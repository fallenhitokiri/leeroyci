package github

import (
	"net/http"
	"testing"

	"github.com/fallenhitokiri/leeroyci/database"
)

func TestAddHeaders(t *testing.T) {
	r, _ := http.NewRequest("GET", "foo", nil)
	addHeaders("foo", r)

	if r.Header["Authorization"][0] != "token foo" {
		t.Error("Wrong authorization headers ", r.Header["Authorization"][0])
	}
}

func TestNewStatusAPI(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("name", "url", "accessKey", false, false)
	j1 := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	j2 := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	com, _ := database.CreateCommand(repo, "", "", "", database.CommandKindBuild)
	database.CreateCommandLog(com, j2, "1", "foo")

	passed := newStatus(j1)
	failed := newStatus(j2)

	if passed.State != statusMessages[statusSuccess]["state"] {
		t.Error("Job did not pass")
	}

	if failed.State != statusMessages[statusFailed]["state"] {
		t.Error("Job did pass")
	}
}

func TestPostStatus(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("name", "url", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	api := &githubMock{}
	postStatus(job, repo, "foo", api)

	if api.method != "POST" {
		t.Error("Wrong method", api.method)
	}

	if api.token != "accessKey" {
		t.Error("Wrong access key", api.token)
	}
}

func TestClosePR(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("name", "url", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")
	api := &githubMock{}
	closePR(job, repo, "foo", api)

	if api.method != "PATCH" {
		t.Error("Wrong method", api.method)
	}

	if api.token != "accessKey" {
		t.Error("Wrong access key", api.token)
	}
}
