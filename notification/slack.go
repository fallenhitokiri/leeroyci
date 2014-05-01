// Implement Slack notifications.
package notification

import (
	"bytes"
	"encoding/json"
	"ironman/config"
	"ironman/logging"
	"log"
	"net/http"
)

type slackPayload struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

// Send a notification to Slack
func slack(c *config.Config, message []byte) {
	_, err := http.Post(
		c.SlackEndpoint,
		"application/json",
		bytes.NewReader(message),
	)

	if err != nil {
		log.Println(err)
	}
}

// Build the payload to send to Slack.
func buildSlack(c *config.Config, job logging.Job) ([]byte, error) {
	payload := slackPayload{
		Channel:  c.SlackChannel,
		Username: "CI",
	}

	message := "Repo: " + job.URL + " Branch: " + job.Branch
	message = message + " Command: " + job.Command + " Pushed by "
	message = message + job.Name + " <" + job.Email + "> "

	if job.Success() == true {
		message = message + "build was successful"
	} else {
		message = message + "build failed"
	}

	payload.Text = message

	marsh, err := json.Marshal(payload)

	if err != nil {
		log.Println(err)
	}

	return marsh, err
}
