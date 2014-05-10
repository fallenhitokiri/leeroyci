package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"ironman/logging"
	"log"
	"net/http"
)

type Comment struct {
	Body string `json:"body"`
}

// Returns a new Comment with the status of the job as body.
func newComment(job logging.Job) Comment {
	c := Comment{}

	if job.Success() {
		c.Body = "build successful"
	} else {
		c.Body = "build failed"
	}

	return c
}

// Post a new comment on a pull request
func PostPR(token string, job logging.Job, pc PRCallback) {
	c := newComment(job)

	m, err := json.Marshal(&c)

	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	r, err := http.NewRequest(
		"POST",
		pc.PR.Comments_url,
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
