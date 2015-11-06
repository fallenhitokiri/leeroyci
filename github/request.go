// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

func postStatus(job *database.Job, repo *database.Repository, URL string) {
	status := newStatus(job)
	payload, err := json.Marshal(&status)

	if err != nil {
		log.Println(err)
		return
	}

	_, err = makeRequest("POST", URL, repo.AccessKey, payload)

	if err != nil {
		log.Println(err)
	}
}

func closePR(job *database.Job, repo *database.Repository, URL string) {
	status := newStatus(job)
	status.State = "closed"
	payload, err := json.Marshal(&status)

	if err != nil {
		log.Println(err)
		return
	}

	_, err = makeRequest("PATCH", URL, repo.AccessKey, payload)

	if err != nil {
		log.Println(err)
	}
}

// githubRequest handles HTTP requests to GitHubs API.
// If the API endpoint does not expect any information nil should be passed as payload.
func makeRequest(method string, url string, token string, payload []byte) ([]byte, error) {
	r, err := http.NewRequest(method, url, bytes.NewReader(payload))

	if err != nil {
		return nil, err
	}

	addHeaders(token, r)

	c := http.Client{}

	re, err := c.Do(r)

	if err != nil {
		return nil, err
	}

	defer re.Body.Close()

	b, err := ioutil.ReadAll(re.Body)

	if err != nil {
		return nil, err
	}

	return b, nil
}

// AddHeaders adds all headers to a request to conform to GitHubs API.
// token is the API token that will be used for the request.
func addHeaders(token string, req *http.Request) {
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "token "+token)
}
