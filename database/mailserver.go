// Package database provides a wrapper between the database and stucts
package database

import (
	"net/mail"
	"net/smtp"
	"strconv"
	"time"
)

// MailServer stores a mailserver configuration.
type MailServer struct {
	ID       int64
	Host     string
	Sender   string
	Port     int
	User     string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// AddMailServer adds a new mail server.
func AddMailServer(host, sender, user, password string, port int) *MailServer {
	m := &MailServer{
		Host:     host,
		Sender:   sender,
		Port:     port,
		User:     user,
		Password: password,
	}

	db.Save(m)

	return m
}

// GetMailServer returns a mail server configuration based on the current configuration.
func GetMailServer() *MailServer {
	m := &MailServer{}
	db.First(m)
	return m
}

// UpdateMailServer updates the existing mail server configuration.
func UpdateMailServer(host, sender, user, password string, port int) *MailServer {
	m := GetMailServer()

	m.Host = host
	m.Sender = sender
	m.User = user
	m.Password = password
	m.Port = port

	db.Save(m)

	return m
}

// DeleteMailServer delete the existing mail server configuration.
func DeleteMailServer() {
	m := GetMailServer()
	db.Delete(m)
}

// Server returns the host name and port for a mailserver.
func (m *MailServer) Server() string {
	return m.Host + ":" + strconv.Itoa(m.Port)
}

// From returns net/mail.Address with sender information for the mail server.
func (m *MailServer) From() mail.Address {
	return mail.Address{
		Name:    "Leeroy CI",
		Address: m.Sender,
	}
}

// Auth returns net/smtp.Auth for this mail server.
func (m *MailServer) Auth() smtp.Auth {
	return smtp.PlainAuth("", m.User, m.Password, m.Host)
}
