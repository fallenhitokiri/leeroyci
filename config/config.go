// Package config takes care of the whole configuration.
package config

import (
	"errors"
	"net/url"
	"strconv"
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
	Users         []User
}

// Repository holds all information needed to identify a repository and run
// tests and builds.
type Repository struct {
	Name      string
	URL       string
	Commands  []Command
	Notify    []Notify
	CommentPR bool
	ClosePR   bool
	StatusPR  bool
	AccessKey string
	Deploy    []Deploy
}

// Command stores a short name and the path or command to execute when a users
// pushes to a repository.
type Command struct {
	Name    string
	Execute string
}

// Notify stores the configuration needed for a notification plugin to work. All
// optiones required by the services are stored as map - it is the job of the
// notification plugin to access them correctly and handle missing ones.
type Notify struct {
	Service   string
	Arguments map[string]string
}

// Deploy stores the command to execute and arguments to pass to the command
// when a users pushes to a certain branch.
type Deploy struct {
	Name      string
	Branch    string
	Execute   string
	Arguments map[string]string
}

// User stores a user account including the password using bcrypt.
type User struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
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

// Identifier returns the name or the URL
func (r *Repository) Identifier() string {
	if r.Name != "" {
		return r.Name
	}
	return r.URL
}

// DeployTarget returns the deployment target for a branch
func (r *Repository) DeployTarget(branch string) (Deploy, error) {
	for _, d := range r.Deploy {
		if d.Branch == branch {
			return d, nil
		}
	}
	return Deploy{}, errors.New("No deployment target for branch")
}
