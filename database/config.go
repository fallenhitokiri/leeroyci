// Package database provides a wrapper between the database and stucts
package database

import (
	"database/sql"
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
