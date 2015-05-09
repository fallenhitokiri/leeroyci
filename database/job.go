// Package database provides a wrapper between the database and stucts
package database

import (
	"time"
)

// Job stores all information about one commit and the executed tasks.
type Job struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Repository   Repository
	RepositoryID int64

	Branch    string
	Commit    string
	CommitURL string

	Name  string
	Email string

	Tasks    []Task `gorm:"many2many:job_tasks;"`
	Deployed *Task
}
