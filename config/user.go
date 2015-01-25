// Package config takes care of the whole configuration.
package config

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User stores a user account including the password using bcrypt.
type User struct {
	Email    string
	Name     string
	Password string
	Admin    bool
}

// GetUser returns the user for a given email address.
func (c *Config) GetUser(email string) (*User, error) {
	for _, u := range c.Users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, errors.New("Could not find user")
}

// CreateUser creates a new user.
func (c *Config) CreateUser(mail, name, pass string, admin bool) (*User, error) {
	_, err := c.GetUser(mail)

	if err == nil {
		return nil, errors.New("User exists")
	}

	hash, err := hashPassword(pass)

	if err != nil {
		return nil, err
	}

	u := &User{
		Email:    mail,
		Name:     name,
		Password: hash,
		Admin:    admin,
	}

	c.mutex.Lock()
	c.Users = append(c.Users, u)
	c.mutex.Unlock()

	err = c.ToFile()

	if err != nil {
		return nil, err
	}

	return u, nil
}

// UpdateUser updates an existing user.
func (c *Config) UpdateUser(mail, name, pass string, admin bool) (*User, error) {
	u, err := c.GetUser(mail)

	if err != nil {
		return nil, err
	}

	hash, err := hashPassword(pass)

	if err != nil {
		return nil, err
	}

	u.Name = name
	u.Password = hash
	u.Admin = admin

	err = c.ToFile()

	if err != nil {
		return nil, err
	}

	return u, err
}

// DeleteUser deletes an existing user.
func (c *Config) DeleteUser(mail string) error {
	c.mutex.Lock()

	i := -1

	// check if the user exists and store the index in the slice
	for c, u := range c.Users {
		if u.Email == mail {
			i = c
		}
	}

	if i == -1 {
		return errors.New("Could not find user.")
	}

	// remove the user
	c.Users[i], c.Users = c.Users[len(c.Users)-1], c.Users[:len(c.Users)-1]

	c.mutex.Unlock()

	err := c.ToFile()

	if err != nil {
		return err
	}

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
