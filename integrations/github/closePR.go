// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"leeroy/logging"
	"log"
	"net/http"
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
		log.Println(err)
		return
	}

	client := &http.Client{}
	r, err := http.NewRequest(
		"PATCH",
		pc.PR.URL,
		bytes.NewReader(m),
	)

	if err != nil {
		log.Println(err)
		return
	}

	addHeaders(token, r)

	_, err = client.Do(r)

	if err != nil {
		log.Println(err)
	}
}

func addHeaders(token string, req *http.Request) {
	t := base64.URLEncoding.EncodeToString([]byte(token))

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Basic "+t)
}
