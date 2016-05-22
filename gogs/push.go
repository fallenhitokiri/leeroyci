package gogs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"
	"github.com/fallenhitokiri/leeroyci/runner"
)

type push struct {
	Secret     string
	Ref        string
	Before     string
	After      string
	CompareURL string `json:"compare_url"`
	Commits    []*commit
	Repository *repository
	Pusher     *user
	Sender     *sender
}

type commit struct {
	ID      string
	Message string
	URL     string
	Author  *user
}

type user struct {
	Name     string
	Email    string
	Username string
}

type repository struct {
	ID            int
	Name          string
	URL           string
	SSHURL        string `json:"ssh_url"`
	CloneURL      string `json:"clone_url"`
	Description   string
	Website       string
	Watchers      int
	Owner         *user
	Private       bool
	DefaultBranch string `json:"default_branch"`
}

type sender struct {
	Login     string
	ID        int
	AvatarURL string `json:"avatar_url"`
}

func (p *push) createJob() error {
	repo := database.GetRepository(p.Repository.URL)

	job := database.CreateJob(
		repo,
		"",
		p.Commits[0].ID,
		p.Commits[0].URL,
		p.Pusher.Name,
		p.Pusher.Email,
	)

	status := make(chan bool, 1)

	queueJob := runner.QueueJob{
		JobID:  job.ID,
		Status: status,
	}

	queueJob.Enqueue()

	return nil
}

func handlePush(req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println(err)
		return
	}

	var callback push
	err = json.Unmarshal(body, &callback)

	if err != nil {
		log.Println("Could not unmarshal request")
		log.Println(err)
		return
	}

	callback.createJob()
}
