// Package database provides a wrapper between the database and stucts
package database

import (
	"net/url"
	"time"
)

// Config represents the complete configuration for the CI.
type Config struct {
	ID       int64
	Secret   string
	URL      string
	Cert     string
	Key      string
	Parallel int

	CreatedAt time.Time
	UpdatedAt time.Time
}

// AddConfig adds a new configuration.
func AddConfig(secret, url, cert, key string, parallel int) *Config {
	c := &Config{
		Secret:   secret,
		URL:      url,
		Cert:     cert,
		Key:      key,
		Parallel: parallel,
	}

	db.Save(c)

	return c
}

// GetConfig returns the current configuration.
func GetConfig() *Config {
	c := &Config{}
	db.First(c)

	if c.Parallel < 1 && c.ID != 0 {
		c.Parallel = 1
		db.Save(c)
	}

	return c
}

// UpdateConfig updates the config.
func UpdateConfig(secret, url, cert, key string, parallel int) *Config {
	c := GetConfig()

	c.Secret = secret
	c.URL = url
	c.Cert = cert
	c.Key = key
	c.Parallel = parallel

	db.Save(c)

	return c
}

// DeleteConfig delete the existing configuration.
func DeleteConfig() {
	c := GetConfig()
	db.Delete(c)
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
