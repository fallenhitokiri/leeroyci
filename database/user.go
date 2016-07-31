package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrorUserNotFound = errors.New("User not found")
var ErrorUserEmailExists = errors.New("User with this Email already exists")

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

// HashPassword generates a hash using bcrypt.
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}
