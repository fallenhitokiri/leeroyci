package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"leeroy/logging"
	"log"
	"net/http"
)

type Update struct {
	State string `json:"state"`
}

// Returns a new Comment with the status of the job as body.
func newUpdate(job *logging.Job) Update {
	u := Update{}
	u.State = "closed"
	return u
}

// Close a pull request if a build failed
func ClosePR(token string, job *logging.Job, pc PRCallback) {
	// just return if the build did not fail
	if job.Success() {
		return
	}

	u := newUpdate(job)

	m, err := json.Marshal(&u)

	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	r, err := http.NewRequest(
		"PATCH",
		pc.PR.Url,
		bytes.NewReader(m),
	)

	if err != nil {
		log.Println(err)
		return
	}

	t := base64.URLEncoding.EncodeToString([]byte(token))

	r.Header.Add("content-type", "application/json")
	r.Header.Add("Authorization", "Basic "+t)

	_, err = client.Do(r)

	if err != nil {
		log.Println(err)
	}
}
