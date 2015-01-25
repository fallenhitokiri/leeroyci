// Package config takes care of the whole configuration.
package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// CreateUserCMD adds a new user to the system.
func (c *Config) CreateUserCMD() {
	m := readEmail()
	n := readName()
	p := readPassword()
	a := readAdmin()

	_, err := c.CreateUser(m, n, p, a)

	if err != nil {
		log.Fatalln("Could not create user: ", err)
	}
}

// UpdateUserCMD updates an existing user.
func (c *Config) UpdateUserCMD() {
	m := readEmail()
	n := readName()
	p := readPassword()
	a := readAdmin()

	_, err := c.UpdateUser(m, n, p, a)

	if err != nil {
		log.Fatalln("Could not update user: ", err)
	}
}

// DeleteUserCMD deletes an existing user.
func (c *Config) DeleteUserCMD() {
	mail := readEmail()
	err := c.DeleteUser(mail)

	if err != nil {
		log.Fatalln("Could not delete user: ", err)
	}
}

// ListUserCMD lists all existing users.
func (c *Config) ListUserCMD() {
	for _, u := range c.Users {
		log.Println("Email: ", u.Email, " Name: ", u.Name, " Admin: ", u.Admin)
	}
}

// ReadEmail reads a users email address from the command line.
func readEmail() string {
	bufReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter email address: ")
	r, err := bufReader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	return strings.TrimRight(r, "\n")
}

// ReadName reads a users name from the command line.
func readName() string {
	bufReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	r, err := bufReader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	return strings.TrimRight(r, "\n")
}

// ReadPassword reads a users password from the command line.
func readPassword() string {
	bufReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter password: ")
	r, err := bufReader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	return strings.TrimRight(r, "\n")
}

// ReadAdmin reads y/n to specify if a user should have admin permissions.
func readAdmin() bool {
	s := bufio.NewScanner(os.Stdin)
	a := false

	fmt.Print("Give admin privileges? [y/n] ")

	for s.Scan() {
		l := s.Text()

		if l == "y" {
			a = true
			break
		} else if l == "n" {
			break
		}

		fmt.Print("\nWrong input, please enter 'y' for yes or 'n' for no: ")
	}

	return a
}
