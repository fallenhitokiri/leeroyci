// Config takes care of the whole configuration.
package config

import (
	"errors"
	"strconv"
)

type Config struct {
	Secret        string
	EmailFrom     string
	EmailHost     string
	EmailPort     int
	EmailUser     string
	EmailPassword string
	SlackChannel  string
	SlackEndpoint string
	Repositories  []Repository
}

type Repository struct {
	URL      string
	Commands []Command
	Notify   []Notify
}

type Command struct {
	Name    string
	Execute string
}

type Notify struct {
	Service   string
	Arguments []string
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

func (c *Config) MailServer() string {
	return c.EmailHost + ":" + strconv.Itoa(c.EmailPort)
}
