package database

// User represents a person using LeeroyCI - this should not be used for a
// service, but only for people who actually login to the web interface.
type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Password  string
	Admin     bool

	APIKeys  []*APIKey
	Sessions []*SessionKey

	db *Database
}

// APIKey stores an API key for a user.
type APIKey struct {
	Key string
}

// SessionKey stores a Session key for a user.
type SessionKey struct {
	Key string
}
