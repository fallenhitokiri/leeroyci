// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"
)

// Payload Slack expects to be POSTed to their API.
type slackPayload struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

// Send a notification to Slack
func sendSlack(job *database.Job, event string) {
	notification, _ := database.GetNotificationForRepoAndType(
		&job.Repository,
		database.NotificationServiceSlack,
	)

	channel, err := notification.GetConfigValue("channel")

	if err != nil {
		log.Print(err)
		return
	}

	endpoint, err := notification.GetConfigValue("endpoint")

	if err != nil {
		log.Print(err)
		return
	}

	payload, err := payloadSlack(job, event, channel)

	_, err = http.Post(
		endpoint,
		"application/json",
		bytes.NewReader(payload),
	)

	if err != nil {
		log.Println(err)
	}
}

// Build the payload to send to Slack.
func payloadSlack(job *database.Job, event, channel string) ([]byte, error) {
	msg := message(job, database.NotificationServiceSlack, event, TypeText)

	payload := slackPayload{
		Channel:  channel,
		Username: "CI",
		Text:     msg,
	}

	marsh, err := json.Marshal(payload)

	if err != nil {
		log.Println(err)
	}

	return marsh, err
}
