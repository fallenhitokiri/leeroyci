package notification

import (
	"net/http"
	"net/http/httptest"
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

func TestSendSlack(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("repo", "", "", false, false)
	job := database.CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	var request *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request = r
	}))
	defer ts.Close()

	database.CreateNotification(
		database.NotificationServiceSlack,
		"channel:::foo:::::endpoint:::"+ts.URL,
		repo,
	)

	sendSlack(job, EventBuild)

	if request.URL.Path != "/" {
		t.Error("Wrong URL path", request.URL.Path)
	}
}
