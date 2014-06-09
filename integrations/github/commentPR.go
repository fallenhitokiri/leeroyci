package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"leeroy/config"
	"leeroy/logging"
	"log"
	"net/http"
)

type Comment struct {
	Body string `json:"body"`
}

// Returns a new Comment with the status of the job as body.
func newComment(job logging.Job, base string) Comment {
	c := Comment{}

	if job.Success() {
		c.Body = "build successful"
	} else {
		c.Body = "build failed - <a href='"
		c.Body = c.Body + base + "status/commit/"
		c.Body = c.Body + job.Hex() + "/" + job.Commit
		c.Body = c.Body + "'>show log</a>"
	}

	return c
}

// Post a new comment on a pull request.
func PostPR(c *config.Config, job logging.Job, pc PRCallback) {
	comment := newComment(job, c.URL)

	m, err := json.Marshal(&comment)

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

	t := base64.URLEncoding.EncodeToString([]byte(c.GitHubKey))

	r.Header.Add("content-type", "application/json")
	r.Header.Add("Authorization", "Basic "+t)

	_, err = client.Do(r)

	if err != nil {
		log.Println(err)
	}
}