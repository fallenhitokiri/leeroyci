// Package database provides a wrapper between the database and stucts
package database

type Task struct {
	ID      int64
	Command string
	Return  string
	Output  string
}
