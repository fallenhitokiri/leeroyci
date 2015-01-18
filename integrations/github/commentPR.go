// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"encoding/json"
	"leeroy/config"
	"leeroy/logging"
	"log"
)

// Everything needed to comment on a GitHub pull request.
type comment struct {
	Body string `json:"body"`
}

// Returns a new Comment with the status of the job as body.
func newComment(job *logging.Job, base string) comment {
	c := comment{}

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

// PostPR posts a new comment on a pull request.
func PostPR(c *config.Config, job *logging.Job, pc PRCallback) {
	comment := newComment(job, c.URL)
	rp, err := c.ConfigForRepo(job.URL)

	if err != nil {
		log.Fatalln(err)
	}

	m, err := json.Marshal(&comment)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = githubRequest("POST", pc.PR.CommentsURL, rp.AccessKey, m)

	if err != nil {
		log.Fatalln(err)
	}
}
