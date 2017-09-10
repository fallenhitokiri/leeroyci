package gitea

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"code.gitea.io/sdk/gitea"
	"github.com/fallenhitokiri/leeroyci/database"
	"github.com/fallenhitokiri/leeroyci/runner"
)

func handlePush(req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println(err)
		return
	}

	var payload gitea.PushPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Println("Could not unmarshal request", err)
		return
	}

	go createJob(payload)
}

func createJob(payload gitea.PushPayload) {
	repo := database.GetRepository(payload.Repo.CloneURL)
	job := database.CreateJob(
		repo,
		payload.Branch(),
		payload.After,
		payload.Commits[0].URL,
		payload.Pusher.FullName,
		payload.Pusher.Email,
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
		// TODO: post status
	}
}
