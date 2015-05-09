// Package database provides a wrapper between the database and stucts
package database

import (
	"database/sql"
	"net/url"
)

// Config represents the complete configuration for the CI.
type Config struct {
	ID           int
	Secret       string
	BuildLogPath string
	URL          string
	Cert         string
	Key          string
	path         string
	Templates    string
	MailServer   MailServer
	MailServerID sql.NullInt64
}

// GetConfig returns the current configuration.
func GetConfig() *Config {
	c := &Config{}
	db.First(c)
	return c
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
