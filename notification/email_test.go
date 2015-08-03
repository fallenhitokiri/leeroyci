package notification

import (
	"testing"
	"time"

	"github.com/fallenhitokiri/leeroyci/database"
)

func TestEmailSubject(t *testing.T) {
	repo := database.Repository{
		Name: "repo",
	}

	job := database.Job{
		Branch:        "branch",
		Commit:        "1234",
		TasksFinished: time.Now(),
		Repository:    repo,
	}

	build := emailSubject(&job, EVENT_BUILD)
	test := emailSubject(&job, EVENT_TEST)
	deployStart := emailSubject(&job, EVENT_DEPLOY_START)
	deployEnd := emailSubject(&job, EVENT_DEPLOY_END)

	if build != "repo/branch build success" {
		t.Error("Wrong message", build)
	}

	if test != "repo/branch tests success" {
		t.Error("Wrong message", test)
	}

	if deployStart != "repo/branch deployment started" {
		t.Error("Wrong message", deployStart)
	}

	if deployEnd != "repo/branch deploy success" {
		t.Error("Wrong message", deployEnd)
	}
}
