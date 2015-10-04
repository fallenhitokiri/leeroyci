// Package database provides a wrapper between the database and stucts
package database

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const sessionDictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const accessKeyDictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const sessionLength = 256
const accessKeyLength = 32

// User stores a user account including the password using bcrypt.
type User struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
	Password  string
	Admin     bool
	Session   string
	AccessKey string
}

// ListUsers returns a list of all users.
func ListUsers() []*User {
	var u []*User
	db.Find(&u)
	return u
}

// GetUser returns the user for a given email address.
func GetUser(email string) (*User, error) {
	u := &User{}
	db.Where("email = ?", email).First(u)

	if u.ID == 0 {
		return nil, errors.New("Could not find user.")
	}

	return u, nil
}

// GetUserBySession returns the user for a given session key.
func GetUserBySession(key string) (*User, error) {
	u := &User{}
	db.Where("session = ?", key).First(u)

	if u.ID == 0 {
		return nil, errors.New("Could not find user.")
	}

	return u, nil
}

// GetUserByAccessKey returns the user for a given access key.
func GetUserByAccessKey(key string) (*User, error) {
	u := &User{}
	db.Where("access_key = ?", key).First(u)

	if u.ID == 0 {
		return nil, errors.New("Could not find user.")
	}

	return u, nil
}

// GetUserByID returns the user for a given ID.
func GetUserByID(id int64) (*User, error) {
	u := &User{}
	db.Where("ID = ?", id).First(u)

	if u.ID == 0 {
		return nil, errors.New("Could not find user.")
	}

	return u, nil
}

// CreateUser creates a new user.
func CreateUser(email, firstName, lastName, password string, admin bool) (*User, error) {
	u, err := GetUser(email)

	if err == nil {
		return nil, errors.New("User with this email address already exists.")
	}

	hash, err := hashPassword(password)

	if err != nil {
		return nil, err
	}

	u = &User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  hash,
		Admin:     admin,
	}

	db.Create(u)

	return u, nil
}

// Update updates an existing user.
func (u *User) Update(email, firstName, lastName, password string, admin bool) (*User, error) {
	u.FirstName = firstName
	u.LastName = lastName
	u.Email = email
	u.Admin = admin

	if password != "" {
		hash, err := hashPassword(password)

		if err != nil {
			return nil, err
		}

		u.Password = hash
	}

	db.Save(u)

	return u, nil
}

// Delete this user.
func (u *User) Delete() error {
	db.Delete(u)

	return nil
}

// NewSession generates a session key and stores it.
func (u *User) NewSession() string {
	for {
		key := generateSessionID(u.Email, sessionDictionary, sessionLength)

		_, err := GetUserBySession(key)

		if err != nil {
			u.Session = key
			db.Save(u)
			return u.Session
		}
	}
}

// NewAccessKey generates a access key and stores it.
func (u *User) NewAccessKey() string {
	for {
		key := generateAccessKey(accessKeyDictionary, accessKeyLength)

		_, err := GetUserByAccessKey(key)

		if err != nil {
			u.AccessKey = key
			db.Save(u)
			return u.AccessKey
		}
	}
}

// generateSessionID generates a new session ID for a user combining the
// email address and a random string.
func generateSessionID(email, dictionary string, length int) string {
	var random = make([]byte, length)
	rand.Read(random)

	for k, v := range random {
		random[k] = dictionary[v%byte(len(dictionary))]
	}

	joined := strings.Join([]string{email, string(random)}, "")

	hash := sha512.New()
	hash.Write([]byte(joined))

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// generateAccessKey generates a new access key for a user.
func generateAccessKey(dictionary string, length int) string {
	var random = make([]byte, length)
	rand.Read(random)

	for k, v := range random {
		random[k] = dictionary[v%byte(len(dictionary))]
	}

	hash := sha512.New()
	hash.Write([]byte(random))

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// HashPassword generates a hash using bcrypt.
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

// ComparePassword returns true if the password matches the hash.
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false
	}

	return true
}
