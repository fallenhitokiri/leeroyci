// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"encoding/json"
	"leeroy/database"
	"log"
	"net/http"
	"time"
)

// TODO: parse full body, not just the fields needed

// PRCallback handles pull requests coming from GitHubs webhook.
type PRCallback struct {
	Number int
	Action string
	PR     PRPullRequest `json:"pull_request"`
}

// PRPullRequest stores the most basic information about a pull request.
type PRPullRequest struct {
	URL         string `json:"url"`
	State       string
	CommentsURL string `json:"comments_url"`
	StatusURL   string `json:"statuses_url"`
	Head        PRCommit
}

// PRCommit points to the latest commit and repository of a pull request.
type PRCommit struct {
	Commit     string `json:"sha"`
	Repository PRRepo `json:"repo"`
}

// PRRepo stores the repository URL of a pull request.
type PRRepo struct {
	HTMLURL string `json:"html_url"`
}

// RepoURL returns the base URL for repository (HTML, not API)
func (p *PRCallback) RepoURL() string {
	return p.PR.Head.Repository.HTMLURL
}

// Handle GitHub pull requests.
func handlePR(req *http.Request) {
	b := parseBody(req)

	var pc PRCallback
	err := json.Unmarshal(b, &pc)

	if err != nil {
		log.Fatalln("Could not unmarshal PR request")
	}

	if pc.Action != "closed" {
		log.Println("handling pull request", pc.Number)
		go updatePR(pc)
	}
}

// Updates the status of a pull request once the build is done. Sleeps 10
// seconds between the checks.
func updatePR(pc PRCallback) {
	counter := 0 // used as pseudo rate limiting so GitHub likes us

	for {
		jobs := database.GetOpenJobs()

		for _, j := range jobs {
			if j.Commit == pc.PR.Head.Commit {
				r := database.RepositoryForURL(j.URL)

				if err != nil {
					log.Fatalln(err)
				}

				if r.CommentPR {
					PostPR(j, pc)
				}

				if r.ClosePR {
					ClosePR(r.AccessKey, j, pc)
				}

				if r.StatusPR {
					if j.Success() {
						PostStatus(
							statusSuccess,
							j.StatusURL(config.CONFIG.URL),
							pc.PR.StatusURL,
							r.AccessKey,
						)
					} else {
						PostStatus(
							statusFailed,
							j.StatusURL(config.CONFIG.URL),
							pc.PR.StatusURL,
							r.AccessKey,
						)
					}
				}

				return
			}
		}

		// Check if the PR is still revelevant or if a new commit was pushed
		// or closed. Terminate the goroutine if this is the case.
		if counter >= 30 {
			if prIsCurrent(pc) == false {
				return
			}
			counter = 0
		} else {
			counter++
		}

		time.Sleep(10 * time.Second)
	}
}

// Returns if PRCallback is for the latest commit.
func prIsCurrent(pc PRCallback) bool {
	rp, err := config.CONFIG.ConfigForRepo(pc.RepoURL())

	r, err := githubRequest("GET", pc.PR.URL, rp.AccessKey, nil)

	if err != nil {
		log.Fatalln(err)
	}

	var pr PRPullRequest
	err = json.Unmarshal(r, &pr)

	if err != nil {
		log.Fatalln(err)
	}

	if pr.Head.Commit != pc.PR.Head.Commit {
		return false
	}

	if pr.State != "open" {
		return false
	}

	return true
}
