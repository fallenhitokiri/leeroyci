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

// Status returns either the exit code of the triggered command or 0 if the
// command finished successfully.
func (j *Job) Status() string {
	code := "0"

	for _, task := range j.Tasks {
		if task.Return != "" {
			return task.Return
		}
	}

	return code
}

// Success returns true if the build was successful.
func (j *Job) Success() bool {
	if j.Status() == "0" {
		return true
	}
	return false
}

// Add a task to the job.
func (j *Job) Add(t Task) {
	j.Tasks = append(j.Tasks, t)
}

// Hex returns the URL in hex
func (j *Job) Hex() string {
	return hex.EncodeToString([]byte(j.URL))
}

// StatusURL returns the URL for the webinterface
func (j *Job) StatusURL(base string) string {
	return fmt.Sprintf("%sstatus/commit/%s/%s", base, j.Hex(), j.Commit)
}

// DeploySuccess returns if the deploy was successful.
func (j *Job) DeploySuccess() bool {
	if j.Deployed == nil {
		return false
	}

	if j.Deployed.Return == "" {
		return true
	}
	return false
}
