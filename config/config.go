// Config takes care of the whole configuration.
package config

import (
	"errors"
	"net/url"
	"strconv"
)

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
}

type Repository struct {
	Name      string
	URL       string
	Commands  []Command
	Notify    []Notify
	CommentPR bool
	ClosePR   bool
	AccessKey string
	Deploy    []Deploy
}

type Command struct {
	Name    string
	Execute string
}

type Notify struct {
	Service   string
	Arguments map[string]string
}

type Deploy struct {
	Name      string
	Branch    string
	Execute   string
	Arguments map[string]string
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

// Retruns the address of the mail server with the port.
func (c *Config) MailServer() string {
	return c.EmailHost + ":" + strconv.Itoa(c.EmailPort)
}

// Returns the URL scheme used.
func (c *Config) Scheme() string {
	u, err := url.Parse(c.URL)

	if err != nil {
		panic(err)
	}

	return u.Scheme
}

// Returns the host.
func (c *Config) Host() string {
	u, err := url.Parse(c.URL)

	if err != nil {
		panic(err)
	}

	return u.Host
}

// Returns the name or the URL
func (r *Repository) Identifier() string {
	if r.Name != "" {
		return r.Name
	}
	return r.URL
}

// Returns the deployment target for a branch
func (r *Repository) DeployTarget(branch string) (Deploy, error) {
	for _, d := range r.Deploy {
		if d.Branch == branch {
			return d, nil
		}
	}
	return Deploy{}, errors.New("No deployment target for branch")
}
