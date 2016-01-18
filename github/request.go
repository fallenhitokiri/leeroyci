// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"
)

var (
	statusSuccess = 1
	statusFailed  = 2
)

// Payload to update / close a PR / commit.
type commitStatus struct {
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

type update struct {
	State string `json:"state"`
}

// newStatus returns a status struct with the correct URL and messages.
func newStatus(job *database.Job) *commitStatus {
	state := statusSuccess

	if !job.Passed() {
		state = statusFailed
	}

	return &commitStatus{
		State:       statusMessages[state]["state"],
		TargetURL:   job.URL(),
		Description: statusMessages[state]["description"],
		Context:     "continuous-integration/leeeroyci",
	}
}

func postStatus(job *database.Job, repo *database.Repository, URL string, api github) {
	status := newStatus(job)
	payload, err := json.Marshal(&status)

	if err != nil {
		log.Println(err)
		return
	}

	_, err = api.makeRequest("POST", URL, repo.AccessKey, payload)

	if err != nil {
		log.Println(err)
	}
}

func closePR(job *database.Job, repo *database.Repository, URL string, api github) {
	status := newStatus(job)
	status.State = "closed"
	payload, err := json.Marshal(&status)

	if err != nil {
		log.Println(err)
		return
	}

	_, err = api.makeRequest("PATCH", URL, repo.AccessKey, payload)

	if err != nil {
		log.Println(err)
	}
}

// AddHeaders adds all headers to a request to conform to GitHubs API.
// token is the API token that will be used for the request.
func addHeaders(token string, req *http.Request) {
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "token "+token)
}
