// Package database provides a wrapper between the database and stucts
package database

import (
	"net/url"
)

// Config represents the complete configuration for the CI.
type Config struct {
	ID     int64
	Secret string
	DBPath string
	URL    string
	Cert   string
	Key    string
}

// AddConfig adds a new configuration.
func AddConfig(secret, dbPath, url, cert, key string) *Config {
	c := &Config{
		Secret: secret,
		DBPath: dbPath,
		URL:    url,
		Cert:   cert,
		Key:    key,
	}

	db.Save(c)

	return c
}

// GetConfig returns the current configuration.
func GetConfig() *Config {
	c := &Config{}
	db.First(c)
	return c
}

// UpdateConfig updates the config.
func UpdateConfig(secret, dbPath, url, cert, key string) *Config {
	c := GetConfig()

	c.Secret = secret
	c.DBPath = dbPath
	c.URL = url
	c.Cert = cert
	c.Key = key

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
