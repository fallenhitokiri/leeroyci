// Package config contains all data models used for LeeroyCI.
package config

// User represents a person using LeeroyCI - this should not be used for a
// service, but only for people who actually login to the web interface.
type User struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
	Admin     bool

	APIKeys []*APIKey
}

// APIKey stores an API key for a user.
type APIKey struct {
	Key string
}
