// Buildlog stores all builds that were triggered and the result.
package logging

import (
	"time"
)

type Buildlog struct {
	Jobs []Job
}

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

// Add adds a new job to the buildlog
func (b *Buildlog) Add(url, branch, command, name, email, output string,
	code error) *Job {
	job := Job{
		URL:        url,
		Branch:     branch,
		Timestamp:  time.Now(),
		Command:    command,
		ReturnCode: code,
		Output:     output,
		Name:       name,
		Email:      email,
	}

	b.Jobs = append(b.Jobs, job)

	return &job
}
