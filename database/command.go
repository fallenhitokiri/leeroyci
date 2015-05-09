// Package database provides a wrapper between the database and stucts
package database

// Command stores a short name and the path or command to execute when a users
// pushes to a repository.
type Command struct {
	ID      int64
	Name    string
	Execute string
}
