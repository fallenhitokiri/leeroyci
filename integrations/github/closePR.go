// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"encoding/json"
	"leeroy/logging"
	"log"
)

// Payload to close a GitHub pull request.
type update struct {
	State string `json:"state"`
}

// Returns a new Comment with the status of the job as body.
func newUpdate() update {
	u := update{}
	u.State = "closed"
	return u
}

// ClosePR closes a pull request if a build failed.
func ClosePR(token string, job *logging.Job, pc PRCallback) {
	// just return if the build did not fail
	if job.Success() {
		return
	}

	u := newUpdate()

	m, err := json.Marshal(&u)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = githubRequest("PATCH", pc.PR.URL, token, m)

	if err != nil {
		log.Fatalln(err)
	}
}
