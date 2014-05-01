// Each triggered build is represented as a Job.
package logging

import (
	"time"
)

type Job struct {
	URL        string
	Branch     string
	Timestamp  time.Time
	Command    string
	ReturnCode error
	Output     string
	Name       string
	Email      string
}

// Returns either the exit code of the triggered command or 0 if the command
// finished successfully.
func (j *Job) Status() string {
	code := "0"

	if j.ReturnCode != nil {
		code = j.ReturnCode.Error()
	}

	return code
}

// Returns true if the build was successful.
func (j *Job) Success() bool {
	if j.Status() == "0" {
		return true
	}
	return false
}
