// Implement Campfire notifications.
package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"leeroy/config"
	"leeroy/logging"
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
func campfire(c *config.Config, j *logging.Job, id string, room string, key string) {
	m, _ := buildCampfire(c, j)

	// Campfire endpoint
	e := fmt.Sprintf(
		"https://%s.campfirenow.com/room/%s/speak.json",
		id,
		room,
	)

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
func buildCampfire(c *config.Config, j *logging.Job) ([]byte, error) {
	m := campfireMessage{}
	p := campfirePayload{Message: &m}

	success := "success"

	if j.Success() == false {
		success = "failed"
	}

	p.Message.Body = fmt.Sprintf(
		"Repo: %s - %s by %s <%s> -> %s\nBuild: %s",
		j.URL,
		j.Branch,
		j.Name,
		j.Email,
		success,
		j.StatusURL(c.URL),
	)

	marsh, err := json.Marshal(p)

	if err != nil {
		log.Println(err)
	}

	return marsh, err
}
