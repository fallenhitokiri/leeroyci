// Package database provides a wrapper between the database and stucts
package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User stores a user account including the password using bcrypt.
type User struct {
	ID        int64
	Email     string
	FirstName string
	LastName  string
	Password  string
	Admin     bool
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
	hash, err := hashPassword(password)

	if err != nil {
		return nil, err
	}

	u.FirstName = firstName
	u.LastName = lastName
	u.Email = email
	u.Password = hash
	u.Admin = admin

	db.Save(u)

	return u, nil
}

// DeleteUser deletes an existing user.
func (u *User) DeleteUser() error {
	db.Delete(u)

	return nil
}

// HashPassword generates a hash using bcrypt.
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

// ComparePassword returns true if the password matches the hash.
func comparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false
	}

	return true
}
