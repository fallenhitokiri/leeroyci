// Package database provides a wrapper between the database and stucts
package database

import (
	"strconv"
)

// MailServer stores a mailserver configuration.
type MailServer struct {
	ID       int
	Host     string
	Sender   string
	Port     int
	User     string
	Password string
}

// GetMailServer returns a MailServer configuration based on the current configuration.
func GetMailServer() *MailServer {
	m := &MailServer{}
	db.First(&m)
	return m
}

// Server returns the host name and port for a mailserver.
func (m *MailServer) Server() string {
	return m.Host + ":" + strconv.Itoa(m.Port)
}
