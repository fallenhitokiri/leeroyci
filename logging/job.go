// Package logging keeps track of all builds and jobs.
package logging

import (
	"encoding/hex"
	"fmt"
	"time"
)

// Job stores all information about one commit and the executed tasks.
type Job struct {
	Identifier string
	URL        string
	Branch     string
	Commit     string
	CommitURL  string
	Timestamp  time.Time
	Name       string
	Email      string
	Tasks      []Task
	Deployed   string
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
	if j.Deployed == "0" {
		return true
	}
	return false
}
