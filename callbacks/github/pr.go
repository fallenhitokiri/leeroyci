// Structs and methods used to process a pull request.
package github

import (
	"encoding/json"
	"ironman/logging"
	"log"
	"net/http"
)

// TODO: parse full body, not just the fields needed

type PRCallback struct {
	Number int
	Action string
	PR     PRPullRequest `json:"pull_request"`
}

type PRPullRequest struct {
	Url          string
	Comments_url string
	Head         PRCommit
}

type PRCommit struct {
	Commit string `json:"sha"`
}

// Handle GitHub pull requests.
func handlePR(req *http.Request, blog *logging.Buildlog) {
	b := parseBody(req)

	var pc PRCallback

	err := json.Unmarshal(b, &pc)
	if err != nil {
		log.Println(string(b))
		panic("Could not unmarshal request")
	}

	go updatePR(pc, blog)
}

func updatePR(pc PRCallback, blog *logging.Buildlog) {
	for {
		<-blog.Done
		for _, job := range blog.Jobs {
			if job.Commit == pc.PR.Head.Commit {
				PostPR("", job, pc)
				ClosePR("", job, pc)
				return
			}
		}
	}
}
