// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"encoding/json"
	"log"
)

// Payload GitHub expects to create a new status.
type status struct {
	State       string `json:"state"`
	TargetURL   string `json:"target_url"`
	Description string `json:"description"`
	Context     string `json:"context"`
}

var (
	statusSuccess = 1
	statusFailed  = 2
	statusPending = 3
)

// status messages linked to their status code.
var statusMessages = map[int]map[string]string{
	statusSuccess: map[string]string{
		"state":       "success",
		"description": "Build successful",
	},
	statusFailed: map[string]string{
		"state":       "failure",
		"description": "Build failed",
	},
	statusPending: map[string]string{
		"state":       "pending",
		"description": "Build pending",
	},
}

// NewStatus returns a status struct with the correct URL and messages.
func newStatus(state int, target string) *status {
	s := status{
		State:       statusMessages[state]["state"],
		TargetURL:   target,
		Description: statusMessages[state]["description"],
		Context:     "continuous-integration/leeeroyci",
	}

	return &s
}

// PostStatus updates the status of a pull request with the build state.
func PostStatus(state int, target, statusURL, accesskey string) {
	s := newStatus(state, target)

	m, err := json.Marshal(&s)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = githubRequest("POST", statusURL, accesskey, m)

	if err != nil {
		log.Fatalln(err)
	}
}
