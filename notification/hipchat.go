// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

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

// Send a notification to HipChat.
func hipchat(n *notification, key string, chl string) {
	e := endpointHipChat(key)
	p := notToHipChapt(n, chl)

	_, err := http.Post(
		e,
		"application/x-www-form-urlencoded",
		bytes.NewReader(p.toURLEncoded()),
	)

	if err != nil {
		log.Println(err)
	}
}

// Convert a notification to a hipchat payload.
func notToHipChapt(n *notification, channel string) hipchatPayload {
	p := hipchatPayload{
		Color:   "green",
		Notify:  true,
		Format:  "text",
		Room:    channel,
		From:    "Leeroy",
		Message: n.rendered,
		Status:  n.Status,
	}

	return p
}

// Build the endpoint for HipChat
func endpointHipChat(key string) string {
	return fmt.Sprintf(
		"https://www.hipchat.com/v1/rooms/message?auth_token=%s",
		key,
	)
}
