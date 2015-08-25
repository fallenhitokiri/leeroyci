package notification

import (
	"testing"

	"github.com/fallenhitokiri/leeroyci/database"
)

func TestEmailSubject(t *testing.T) {
	repo, _ := database.CreateRepository("repo", "bar", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "1234", "commitURL", "foo", "bar")
	job.TasksDone()

	build := emailSubject(job, EVENT_BUILD)
	test := emailSubject(job, EVENT_TEST)
	deployStart := emailSubject(job, EVENT_DEPLOY_START)
	deployEnd := emailSubject(job, EVENT_DEPLOY_END)

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
