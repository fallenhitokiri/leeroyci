// Structs and methods used to process a pull request.
package github

import (
	"encoding/json"
	"ironman/config"
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
func handlePR(req *http.Request, blog *logging.Buildlog, c *config.Config) {
	b := parseBody(req)

	var pc PRCallback

	log.Println("handling pull request", pc.Number)

	err := json.Unmarshal(b, &pc)
	if err != nil {
		log.Println(string(b))
		panic("Could not unmarshal request")
	}

	go updatePR(pc, blog, c)
}

func updatePR(pc PRCallback, blog *logging.Buildlog, c *config.Config) {
	for {
		for _, j := range blog.Jobs {
			if j.Commit == pc.PR.Head.Commit {
				r, err := c.ConfigForRepo(j.URL)

				if err != nil {
					log.Println(err)
					return
				}

				if r.CommentPR {
					PostPR(c.GitHubKey, j, pc)
				}

				if r.ClosePR {
					ClosePR(c.GitHubKey, j, pc)
				}

				return
			}
		}
		<-blog.Done
	}
}
