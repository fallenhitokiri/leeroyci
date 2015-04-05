// Package config takes care of the whole configuration.
package config

import (
	"errors"
	"net/url"
	"strconv"
	"sync"
)

// Config represents the complete configuration for the CI.
type Config struct {
	Secret        string
	BuildLogPath  string
	EmailFrom     string
	EmailHost     string
	EmailPort     int
	EmailUser     string
	EmailPassword string
	Repositories  []Repository
	URL           string
	Cert          string
	Key           string
	path          string
	Templates     string
	Users         []*User
	mutex         sync.Mutex
}

// ConfigForRepo returns the configuration for a repository that matches
// the URL.
func (c *Config) ConfigForRepo(url string) (Repository, error) {
	r := Repository{}

	for _, repo := range c.Repositories {
		if repo.URL == url {
			r = repo
			return r, nil
		}
	}

	msg := "Could not find repository with URL: " + url
	err := errors.New(msg)
	return r, err
}

// MailServer retruns the address of the mail server with the port.
func (c *Config) MailServer() string {
	return c.EmailHost + ":" + strconv.Itoa(c.EmailPort)
}

// Scheme returns the URL scheme used.
func (c *Config) Scheme() string {
	u, err := url.Parse(c.URL)

	if err != nil {
		panic(err)
	}

	return u.Scheme
}

// Host returns the host.
func (c *Config) Host() string {
	u, err := url.Parse(c.URL)

	if err != nil {
		panic(err)
	}

	return u.Host
}
