// Package config contains all data models used for LeeroyCI.
package config

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"

	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	sessionDictionary   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	accessKeyDictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_?!+=%$&/()"
	sessionLength       = 256
	accessKeyLength     = 64

	errorMissingValue = "Email and password are required"
	errorUserExists   = "User with this email address already exists."
)

// User represents a person using LeeroyCI - this should not be used for a
// service, but only for people who actually login to the web interface.
type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Password  string
	Admin     bool

	APIKeys  []*APIKey
	Sessions []*SessionKey
}

// APIKey stores an API key for a user.
type APIKey struct {
	Key string
}

// SessionKey stores a Session key for a user.
type SessionKey struct {
	Key string
}

// ListUsers returns all users known to the system.
func ListUsers() []*User {
	return cfg.Users
}

// GetUserByEmail returns a user for a given email address.
func GetUserByEmail(email string) (*User, error) {
	email = strings.ToLower(email)
	for _, user := range cfg.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("User not found")
}

// GetUserBySessionKey returns a user for a given session key.
func GetUserBySessionKey(key string) (*User, error) {
	for _, user := range cfg.Users {
		for _, session := range user.Sessions {
			if session.Key == key {
				return user, nil
			}
		}
	}
	return nil, errors.New("User not found")
}

// GetUserByAPIKey returns a user for a given API key.
func GetUserByAPIKey(key string) (*User, error) {
	for _, user := range cfg.Users {
		for _, apiKey := range user.APIKeys {
			if apiKey.Key == key {
				return user, nil
			}
		}
	}
	return nil, errors.New("User not found")
}

// GetUserByID returns a user for a given ID.
func GetUserByID(id string) (*User, error) {
	for _, user := range cfg.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("User not found")
}

// CreateUser checks if the user already exists, if this is not the case it
// generates a new UUID and adds the user to the config.
func CreateUser(user *User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New(errorMissingValue)
	}

	if _, err := GetUserByEmail(user.Email); err == nil {
		return errors.New(errorUserExists)
	}

	uuid, err := newUUID()
	if err != nil {
		return err
	}
	user.ID = uuid

	pass, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pass

	cfg.Users = append(cfg.Users, user)
	return cfg.Save()
}

// UpdatePassword updates a users password.
func (u *User) UpdatePassword(password string) error {
	pass, err := hashPassword(password)

	if err != nil {
		return err
	}

	u.Password = pass

	return cfg.Save()
}

// Delete deletes a user.
func (u *User) Delete() error {
	for index, user := range cfg.Users {
		if user.ID == u.ID {
			cfg.Users = append(cfg.Users[:index], cfg.Users[index+1:]...)
			break
		}
	}
	return cfg.Save()
}

// NewSession generates a session key and stores it.
func (u *User) NewSession() (string, error) {
	for {
		key := generateSessionID(u.Email, sessionDictionary, sessionLength)

		_, err := GetUserBySessionKey(key)

		if err != nil {
			session := &SessionKey{
				Key: key,
			}
			u.Sessions = append(u.Sessions, session)
			return key, cfg.Save()
		}
	}
}

// NewAccessKey generates a access key and stores it.
func (u *User) NewAccessKey() (string, error) {
	for {
		key := generateAccessKey(accessKeyDictionary, accessKeyLength)

		_, err := GetUserByAPIKey(key)

		if err != nil {
			api := &APIKey{
				Key: key,
			}
			u.APIKeys = append(u.APIKeys, api)
			return key, cfg.Save()
		}
	}
}

// CheckPassword returns true if the password is valid.
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// newUUID generates a new UUID - for every generated ID check if a user with
// this ID already exists to prevent resuing an ID.
func newUUID() (string, error) {
	for {
		id, err := uuid.NewV4()

		if err != nil {
			return "", err
		}

		_, err = GetUserByID(id.String())

		if err != nil {
			return id.String(), nil
		}
	}
}

// hashPassword generates a hash using bcrypt.
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
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

	return string(random)
}
