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
	State       string
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
		time.Sleep(10 * time.Second)

		job := database.GetJobByCommit(p.PR.Head.Commit)

		log.Println("Handling", job)

		if job.ID == 0 {
			log.Println("job does not exist")
			continue
		}

		nilTime := time.Time{}
		if job.TasksFinished == nilTime {
			log.Println("job not finished")
			continue
		}

		log.Println("Updating PR")
		return
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

	go callback.updatePR()
}
