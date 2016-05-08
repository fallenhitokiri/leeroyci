package data

import "time"

// Result keeps track of the different tasks that ran for a specific commit.
type Result struct {
	Commit    string
	CommitURL string
	Cancelled bool

	Start  time.Time
	End    time.Time
	Passed bool

	TaskResults []*TaskResult
}
