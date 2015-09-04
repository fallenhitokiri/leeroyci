// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fallenhitokiri/leeroyci/database"
)

var (
	statusSuccess = 1
	statusFailed  = 2
)

type pullRequestCallback struct {
	Number int
	Action string
	PR     pullRequest `json:"pull_request"`
}

type pullRequest struct {
	URL         string `json:"url"`
	State       string `json:"state"`
	CommentsURL string `json:"comments_url"`
	StatusURL   string `json:"statuses_url"`
	Head        pullRequestCommit
}

type pullRequestCommit struct {
	Commit     string                `json:"sha"`
	Repository pullRequestRepository `json:"repo"`
}

type pullRequestRepository struct {
	HTMLURL string `json:"html_url"`
}

// Payload to update / close a PR / commit.
type postStatus struct {
	State       string `json:"state"`
	TargetURL   string `json:"target_url"`
	Description string `json:"description"`
	Context     string `json:"context"`
}

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
}

func (p *pullRequestCallback) repositoryURL() string {
	return p.PR.Head.Repository.HTMLURL
}

func (p *pullRequestCallback) updatePR() {
	for {
		if p.isCurrent() == false {
			log.Println("not current")
			return
		}

		job := database.GetJobByCommit(p.PR.Head.Commit)

		if job.ID == 0 {
			time.Sleep(30 * time.Second)
			continue
		}

		nilTime := time.Time{}
		if !job.TasksFinished.After(nilTime) {
			time.Sleep(10 * time.Second)
			continue
		}

		repository, err := database.GetRepositoryByID(job.RepositoryID)

		if err != nil {
			log.Println(err)
			return
		}

		if repository.StatusPR {
			p.postStatus(job, repository)
		}

		if repository.ClosePR && job.Passed() == false {
			p.closePR(job, repository)
		}

		return
	}
}

func (p *pullRequestCallback) isCurrent() bool {
	repo := database.GetRepository(p.repositoryURL())
	response, err := makeRequest("GET", p.PR.URL, repo.AccessKey, nil)

	log.Println("----------------------------------------------")
	log.Println("Accesskey", repo.AccessKey)
	log.Println("URL", p.PR.URL)

	if err != nil {
		log.Println("err1")
		log.Println(err)
		log.Println("----------------------------------------------")
		return false
	}

	var pr pullRequest
	err = json.Unmarshal(response, &pr)

	if err != nil {
		log.Println("Unmarshal")
		log.Println(err)
		log.Println("----------------------------------------------")
		return false
	}

	if pr.Head.Commit != p.PR.Head.Commit {
		log.Println("Head fetched", pr.Head.Commit)
		log.Println("PR Head", p.PR.Head.Commit)
		log.Println("----------------------------------------------")
		return false
	}

	if pr.State != "open" {
		log.Println("not open")
		log.Println("----------------------------------------------")
		return false
	}

	return true
}

func (p *pullRequestCallback) postStatus(job *database.Job, repo *database.Repository) {
	status := newStatus(job)
	payload, err := json.Marshal(&status)

	if err != nil {
		log.Println(err)
		return
	}

	_, err = makeRequest("POST", p.PR.StatusURL, repo.AccessKey, payload)

	if err != nil {
		log.Println(err)
	}
}

type update struct {
	State string `json:"state"`
}

func (p *pullRequestCallback) closePR(job *database.Job, repo *database.Repository) {
	status := newStatus(job)
	status.State = "closed"
	payload, err := json.Marshal(&status)

	if err != nil {
		log.Println(err)
		return
	}

	_, err = makeRequest("PATCH", p.PR.URL, repo.AccessKey, payload)

	if err != nil {
		log.Println(err)
	}
}

// newStatus returns a status struct with the correct URL and messages.
func newStatus(job *database.Job) *postStatus {
	state := statusSuccess

	if !job.Passed() {
		state = statusFailed
	}

	return &postStatus{
		State:       statusMessages[state]["state"],
		TargetURL:   job.URL(),
		Description: statusMessages[state]["description"],
		Context:     "continuous-integration/leeeroyci",
	}
}

func handlePR(req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println(err)
		return
	}

	var callback pullRequestCallback

	err = json.Unmarshal(body, &callback)

	if err != nil {
		log.Println("Could not unmarshal request")
		return
	}

	if callback.Action != "closed" {
		go callback.updatePR()
	}
}
