package notification

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fallenhitokiri/leeroyci/database"
)

func TestEndpointCampfire(t *testing.T) {
	end := endpointCampfire("1", "2")

	if end != "https://1.campfirenow.com/room/2/speak.json" {
		t.Error("Wrong endpoint", end)
	}
}

func TestPayloadCampfire(t *testing.T) {
	repo, _ := database.CreateRepository("repo", "", "", false, false)
	job := database.CreateJob(repo, "branch", "bar", "commit URL", "name", "email")

	pay, err := payloadCampfire(job, EventBuild)

	if err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(string(pay), "repo") {
		t.Error("Wrong payload", string(pay))
	}
}

func TestRequestCampfire(t *testing.T) {
	r := requestCampfire("foo", "bar", []byte("baz"))

	if r.Method != "POST" {
		t.Error("Wrong method ", r.Method)
	}

	u, _, _ := r.BasicAuth()

	if u != "bar" {
		t.Error("Wrong username", u)
	}
}

func TestSendCampfire(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("repo", "", "", false, false)
	job := database.CreateJob(repo, "branch", "bar", "commit URL", "name", "email")
	database.CreateNotification(
		database.NotificationServiceCampfire,
		"id:::foo:::::room:::bar:::::key:::baz",
		repo,
	)

	var request *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request = r
	}))
	defer ts.Close()

	campfireEndpoint = ts.URL + "/%s/%s"

	sendCampfire(job, EventBuild)

	if request.URL.Path != "/foo/bar" {
		t.Error("Wrong URL path", request.URL.Path)
	}

	if request.Header["Authorization"][0] != "Basic YmF6Olg=" {
		t.Error("Wrong auth token", request.Header["Authorization"][0])
	}
}
