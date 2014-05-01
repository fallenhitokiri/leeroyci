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
