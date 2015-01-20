// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Payload Campfire expects to be POSTed to their API.
type campfirePayload struct {
	Message *campfireMessage `json:"message"`
}

// Message part of the payload Campfire expects.
type campfireMessage struct {
	Body string `json:"body"`
}

// Send a notification to Campfire
func campfire(n *notification, id string, room string, key string) {
	m, _ := buildCampfire(n)
	e := endpointCampfire(id, room)
	r := requestCampfire(e, key, m)
	c := &http.Client{}

	_, err := c.Do(r)

	if err != nil {
		log.Fatal(err)
	}
}

// Build the payload to send to Campfire.
func buildCampfire(n *notification) ([]byte, error) {
	p := campfirePayload{
		Message: &campfireMessage{
			Body: n.message,
		},
	}

	marsh, err := json.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}

	return marsh, err
}

// Build the endpoint for campfire
func endpointCampfire(id, room string) string {
	return fmt.Sprintf(
		"https://%s.campfirenow.com/room/%s/speak.json",
		id,
		room,
	)
}

// Build the request for the campfire API.
func requestCampfire(endpoint string, key string, message []byte) *http.Request {
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(message))

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
