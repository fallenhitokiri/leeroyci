// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"
)

var campfireEndpoint = "https://%s.campfirenow.com/room/%s/speak.json"

// Payload Campfire expects to be POSTed to their API.
type campfirePayload struct {
	Message *campfireMessage `json:"message"`
}

// Message part of the payload Campfire expects.
type campfireMessage struct {
	Body string `json:"body"`
}

// sendCampfire sends a notification to Campfire
func sendCampfire(job *database.Job, event string) {
	payload, err := payloadCampfire(job, event)

	if err != nil {
		log.Println(err)
		return
	}

	not, _ := database.GetNotificationForRepoAndType(
		&job.Repository,
		database.NotificationServiceCampfire,
	)

	id, err := not.GetConfigValue("id")

	if err != nil {
		log.Println(err)
		return
	}

	room, err := not.GetConfigValue("room")

	if err != nil {
		log.Println(err)
		return
	}

	endpoint := endpointCampfire(id, room)

	key, err := not.GetConfigValue("key")

	if err != nil {
		log.Println(err)
		return
	}

	request := requestCampfire(endpoint, key, payload)
	client := &http.Client{}

	_, err = client.Do(request)

	if err != nil {
		log.Println(err)
	}
}

// Build the payload to send to Campfire.
func payloadCampfire(job *database.Job, event string) ([]byte, error) {
	msg := message(job, database.NotificationServiceCampfire, event, TypeText)

	p := campfirePayload{
		Message: &campfireMessage{
			Body: msg,
		},
	}

	return json.Marshal(p)
}

// Build the endpoint for campfire
func endpointCampfire(id, room string) string {
	return fmt.Sprintf(campfireEndpoint, id, room)
}

// Build the request for the campfire API.
func requestCampfire(endpoint string, key string, payload []byte) *http.Request {
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(payload))

	if err != nil {
		log.Fatal(err)
	}

	// There is no need for a password. Campire API documentation suggests
	// to use X so a password is present in case a component of the
	// implementation has problems without one.
	req.SetBasicAuth(key, "X")
	req.Header.Add("Content-Type", "application/json")

	return req
}
