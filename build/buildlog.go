// Buildlog stores all builds that were triggered and the result.
package build

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
	ReturnCode error
	Output     string
	Pusher     string
}

// Add adds a new job to the buildlog
func (b *Buildlog) Add(url, branch, pusher, output string, code error) *Job {
	job := Job{
		URL:        url,
		Branch:     branch,
		Timestamp:  time.Now(),
		ReturnCode: code,
		Output:     output,
		Pusher:     pusher,
	}

	b.Jobs = append(b.Jobs, job)

	return &job
}
