// Package config takes care of the whole configuration.
package config

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// User stores a user account including the password using bcrypt.
type User struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
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

// AddUser adds a new user to the system.
func (c *Config) AddUser(reader io.Reader) {
	mail, first, last, pass := readUser(reader)
	_, err := c.GetUser(mail)

	if err == nil {
		log.Fatalln("User exists")
	}

	hash, err := hashPassword(pass)

	if err != nil {
		log.Fatalln(err)
	}

	u := &User{
		Email:     mail,
		FirstName: first,
		LastName:  last,
		Password:  hash,
	}

	c.Users = append(c.Users, u)

	err = c.ToFile()

	if err != nil {
		log.Fatalln(err)
	}
}

// ReadUser reads user information from the command line.
func readUser(reader io.Reader) (string, string, string, string) {
	if reader == nil {
		reader = os.Stdin
	}

	bufReader := bufio.NewReader(reader)

	fmt.Print("Enter email address: ")
	e, err := bufReader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Enter first name: ")
	f, err := bufReader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Enter last name: ")
	l, err := bufReader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Enter password: ")
	p, err := bufReader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	e = strings.TrimRight(e, "\n")
	f = strings.TrimRight(f, "\n")
	l = strings.TrimRight(l, "\n")
	p = strings.TrimRight(p, "\n")

	return e, f, l, p
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
