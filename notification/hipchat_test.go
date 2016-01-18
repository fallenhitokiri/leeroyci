// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fallenhitokiri/leeroyci/database"
)

func TestBuildHipChat(t *testing.T) {
	repo := database.Repository{
		Name: "repo",
	}

	j := database.Job{
		Repository: repo,
		Branch:     "branch",
		Name:       "name",
		Email:      "email",
	}

	p := payloadHipchat(&j, "foo", "bar")

	if p.Room != "bar" {
		t.Error("Wrong room", p.Room)
	}
}

func TestToURLEncoded(t *testing.T) {
	h := hipchatPayload{
		Room:    "foo",
		From:    "bar",
		Message: "baz",
		Notify:  true,
		Format:  "text",
		Status:  true,
	}

	e := string(h.toURLEncoded())

	if strings.Contains(e, "notify=1") == false {
		t.Error("Wrong notification settings")
	}

	if strings.Contains(e, "color=green") == false {
		t.Error("Wrong notification color")
	}

	h.Status = false
	h.Notify = false

	e = string(h.toURLEncoded())

	if strings.Contains(e, "notify=2") == false {
		t.Error("Wrong notification settings")
	}

	if strings.Contains(e, "color=red") == false {
		t.Error("Wrong notification color")
	}
}

func TestEndpointHipChat(t *testing.T) {
	exp := "https://www.hipchat.com/v1/rooms/message?auth_token=foo"
	e := endpointHipChat("foo")

	if e != exp {
		t.Error("Wrong API endpoint ", e)
	}
}

func TestSendHipchat(t *testing.T) {
	database.NewInMemoryDatabase()
	repo, _ := database.CreateRepository("repo", "", "", false, false)
	job := database.CreateJob(repo, "branch", "bar", "commit URL", "name", "email")
	database.CreateNotification(
		database.NotificationServiceHipchat,
		"channel:::foo:::::key:::bar",
		repo,
	)

	var request *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request = r
	}))
	defer ts.Close()

	hipchatEndpoint = ts.URL + "/%s"

	sendHipchat(job, EventBuild)

	if request.URL.Path != "/bar" {
		t.Error("Wrong URL path", request.URL.Path)
	}
}
