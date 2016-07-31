package database

import "time"

// TaskResult keeps track of the results of a finished task.
type TaskResult struct {
	ID   string
	Name string
	Kind string

	Start  time.Time
	End    time.Time
	Passed bool

	ReturnCode string
	Output     string
}
