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

		if repository.ClosePR && job.Passed() == false {
			closePR(job, repository, p.PR.URL, githubAPI{})
		}

		return
	}
}

func (p *pullRequestCallback) isCurrent() bool {
	repo := database.GetRepository(p.repositoryURL())
	response, err := githubAPI{}.makeRequest("GET", p.PR.URL, repo.AccessKey, nil)

	if err != nil {
		return false
	}

	var pr pullRequest
	err = json.Unmarshal(response, &pr)

	if err != nil {
		return false
	}

	if pr.Head.Commit != p.PR.Head.Commit {
		return false
	}

	if pr.State != "open" {
		return false
	}

	return true
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
