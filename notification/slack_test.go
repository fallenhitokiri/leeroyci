package notification

import (
	"testing"

	"github.com/fallenhitokiri/leeroyci/database"
)

func TestBuildSlack(t *testing.T) {
	repo := database.Repository{
		Name: "repo",
	}

	job := database.Job{
		Branch:     "branch",
		Commit:     "1234",
		Repository: repo,
	}

	_, err := payloadSlack(&job, EventBuild, "foo")

	if err != nil {
		t.Error(err)
	}
}
