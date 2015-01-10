// Implement Slack notifications.
package notification

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type slackPayload struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

// Send a notification to Slack
func slack(n *notification, endpoint string, channel string) {
	m, err := buildSlack(n, channel)

	_, err = http.Post(
		endpoint,
		"application/json",
		bytes.NewReader(m),
	)

	if err != nil {
		log.Println(err)
	}
}

// Build the payload to send to Slack.
func buildSlack(n *notification, channel string) ([]byte, error) {
	p := slackPayload{
		Channel:  channel,
		Username: "CI",
		Text:     n.rendered,
	}

	marsh, err := json.Marshal(p)

	if err != nil {
		log.Println(err)
	}

	return marsh, err
}
