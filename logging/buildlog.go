// Buildlog stores all builds that were triggered and the result.
package logging

import (
	"time"
)

type Buildlog struct {
	Jobs []Job
}

// Add adds a new job to the buildlog
func (b *Buildlog) Add(url, branch, commit, command, name, email, output string,
	code error) *Job {
	job := Job{
		URL:        url,
		Branch:     branch,
		Timestamp:  time.Now(),
		Commit:     commit,
		Command:    command,
		ReturnCode: code,
		Output:     output,
		Name:       name,
		Email:      email,
	}

	b.Jobs = append(b.Jobs, job)

	return &job
}
