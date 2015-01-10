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

type campfirePayload struct {
	Message *campfireMessage `json:"message"`
}

type campfireMessage struct {
	Body string `json:"body"`
}

// Send a notification to Campfire
func campfire(n *notification, id string, room string, key string) {
	m, _ := buildCampfire(n)
	e := endpointCampfire(id, room)

	client := &http.Client{}

	req, err := http.NewRequest("POST", e, bytes.NewReader(m))

	// There is no need for a password. Campire API documentation suggests
	// to use X so a password is present in case a component of the
	// implementation has problems without one.
	req.SetBasicAuth(key, "X")
	req.Header.Add("Content-Type", "application/json")

	client.Do(req)

	if err != nil {
		log.Println(err)
	}
}

// Build the payload to send to Campfire.
func buildCampfire(n *notification) ([]byte, error) {
	p := campfirePayload{
		Message: &campfireMessage{
			Body: n.rendered,
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
