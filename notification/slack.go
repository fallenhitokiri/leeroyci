// Implement Slack notifications.
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

type slackPayload struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

// Send a notification to Slack
func slack(c *config.Config, j *logging.Job, ep string, chl string) {
	m, err := buildSlack(c, j, chl)

	_, err = http.Post(
		ep,
		"application/json",
		bytes.NewReader(m),
	)

	if err != nil {
		log.Println(err)
	}
}

// Build the payload to send to Slack.
func buildSlack(c *config.Config, j *logging.Job, chl string) ([]byte, error) {
	p := slackPayload{
		Channel:  chl,
		Username: "CI",
	}

	success := "success"

	if j.Success() == false {
		success = "failed"
	}

	m := fmt.Sprintf(
		"Repo: %s - %s by %s <%s> -> %s\nBuild: %s",
		j.URL,
		j.Branch,
		j.Name,
		j.Email,
		success,
		j.StatusURL(c.URL),
	)

	p.Text = m

	marsh, err := json.Marshal(p)

	if err != nil {
		log.Println(err)
	}

	return marsh, err
}
