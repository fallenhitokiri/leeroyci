// Package database provides a wrapper between the database and stucts
package database

// Notify stores the configuration needed for a notification plugin to work. All
// optiones required by the services are stored as map - it is the job of the
// notification plugin to access them correctly and handle missing ones.
type Notify struct {
	ID        int64
	Service   string
	Arguments string
}
