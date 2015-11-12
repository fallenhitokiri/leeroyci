// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/fallenhitokiri/leeroyci/database"
)

var hipchatEndpoint = "https://www.hipchat.com/v1/rooms/message?auth_token=%s"

// Payload HipChat expects to be POSTed to the API.
type hipchatPayload struct {
	Room    string
	From    string
	Color   string
	Message string
	Notify  bool
	Format  string
	Status  bool
}

// HipChat expects www-form-urlencoded - prepare the struct.
func (h *hipchatPayload) toURLEncoded() []byte {
	d := url.Values{}
	d.Add("room_id", h.Room)
	d.Add("from", h.From)
	d.Add("message", h.Message)
	d.Add("message_format", h.Format)

	if h.Notify == true {
		d.Add("notify", "1")
	} else {
		d.Add("notify", "2")
	}

	if h.Status == true {
		d.Add("color", "green")
	} else {
		d.Add("color", "red")
	}

	return []byte(d.Encode())
}

func sendHipchat(job *database.Job, event string) {
	not, _ := database.GetNotificationForRepoAndType(
		&job.Repository,
		database.NotificationServiceHipchat,
	)

	channel, err := not.GetConfigValue("channel")

	if err != nil {
		log.Println(err)
		return
	}

	payload := payloadHipchat(job, event, channel)

	key, err := not.GetConfigValue("key")

	if err != nil {
		log.Println(err)
		return
	}

	endpoint := endpointHipChat(key)

	_, err = http.Post(
		endpoint,
		"application/x-www-form-urlencoded",
		bytes.NewReader(payload.toURLEncoded()),
	)

	if err != nil {
		log.Println(err)
	}
}

// Convert a job to a hipchat payload.
func payloadHipchat(job *database.Job, event, channel string) hipchatPayload {
	msg := message(job, database.NotificationServiceHipchat, event, TypeText)

	return hipchatPayload{
		Color:   "green",
		Notify:  true,
		Format:  "text",
		Room:    channel,
		From:    "LeeroyCI",
		Message: msg,
		Status:  job.Passed(),
	}
}

// Build the endpoint for HipChat
func endpointHipChat(key string) string {
	return fmt.Sprintf(hipchatEndpoint, key)
}
