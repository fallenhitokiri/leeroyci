// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fallenhitokiri/leeroyci/database"
	"github.com/fallenhitokiri/leeroyci/runner"
)

type pushCallback struct {
	Ref        string
	After      string
	Before     string
	Created    bool
	Deleted    bool
	Forced     bool
	Compare    string
	Commits    []pushCommit
	HeadCommit pushCommit `json:"head_commit"`
	Repository pushRepository
	Pusher     pushUser
}

type pushCommit struct {
	ID        string
	Distinct  bool
	Message   string
	Timestamp string
	URL       string
	Author    pushUser
	Committer pushUser
	Added     []string
	Removed   []string
	Modified  []string
}

type pushUser struct {
	Name  string
	Email string
}

type pushRepository struct {
	ID          int64
	Name        string
	URL         string
	Description string
	CreatedAt   int64  `json:"created_at"`
	PushedAt    int64  `json:"pushed_at"`
	StatusURL   string `json:"statuses_url"`
}

// repositoryURL returns the URL for the repository
func (p *pushCallback) repositoryURL() string {
	return p.Repository.URL
}

// branch returns the name of the branch.
func (p *pushCallback) branch() string {
	s := strings.Split(p.Ref, "/")
	return s[2]
}

// returns the ID of the head commit.
func (p *pushCallback) commit() string {
	return p.HeadCommit.ID
}

// commitURL returns the URL to the head commit.
func (p *pushCallback) commitURL() string {
	return p.HeadCommit.URL
}

// name returns the name of the git user.
func (p *pushCallback) name() string {
	return p.Pusher.Name
}

// email returns the email of the git user.
func (p *pushCallback) email() string {
	return p.Pusher.Email
}

// shouldRun returns if this push should create a job.
func (p *pushCallback) shouldRun() bool {
	if p.Deleted == true {
		return false
	}
	return true
}

func (p *pushCallback) statusURL() string {
	return strings.Replace(p.Repository.StatusURL, "{sha}", p.commit(), 1)
}

// createJob adds a new job to the database.
func (p *pushCallback) createJob() error {
	if p.shouldRun() == false {
		log.Println("Not adding", p.repositoryURL(), p.branch())
		return nil
	}

	repo := database.GetRepository(p.repositoryURL())

	job := database.CreateJob(
		repo,
		p.branch(),
		p.commit(),
		p.commitURL(),
		p.name(),
		p.email(),
	)

	status := make(chan bool, 1)

	queueJob := runner.QueueJob{
		JobID:  job.ID,
		Status: status,
	}

	queueJob.Enqueue()

	if repo.StatusPR {
		<-status
		job = database.GetJob(job.ID)
		postStatus(job, repo, p.statusURL(), githubAPI{})
	}

	return nil
}

func handlePush(req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println(err)
		return
	}

	var callback pushCallback

	err = json.Unmarshal(body, &callback)

	if err != nil {
		log.Println("Could not unmarshal request")
		return
	}

	go callback.createJob()
}
