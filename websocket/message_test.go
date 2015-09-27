package websocket

import (
	"testing"

	"github.com/fallenhitokiri/leeroyci/database"
)

func TestNewMessage(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("foo", "baz", "accessKey", false, false)
	job := database.CreateJob(repo, "branch", "commit", "commitURL", "name", "email")

	msg := NewMessage(job, "foo")

	if msg.Status != database.JobStatusPending {
		t.Error("Wrong status", msg.Status)
	}

	if msg.RepositoryName != "foo" {
		t.Error("Wrong repository name", msg.RepositoryName)
	}
}
