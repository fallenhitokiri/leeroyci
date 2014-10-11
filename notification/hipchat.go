// Implement HipChat notifications.
package notification

import (
	"bytes"
	"fmt"
	"leeroy/config"
	"leeroy/logging"
	"log"
	"net/http"
	"net/url"
)

var api = "https://www.hipchat.com/v1/rooms/message?auth_token=%s"

type hipchatPayload struct {
	Room    string
	From    string
	Color   string
	Message string
	Notify  bool
	Format  string
}

// HipChat expects www-form-urlencoded - prepare the struct.
func (h *hipchatPayload) toURLEncoded() []byte {
	d := url.Values{}
	d.Add("room_id", h.Room)
	d.Add("from", h.From)
	d.Add("message", h.Message)
	d.Add("message_format", h.Format)
	d.Add("color", h.Color)

	if h.Notify == true {
		d.Add("notify", "1")
	} else {
		d.Add("notify", "2")
	}

	return []byte(d.Encode())
}

func hipchat(c *config.Config, j *logging.Job, key string, chl string) {
	endpoint := fmt.Sprintf(api, key)

	p := buildHipChat(c, j, chl)

	_, err := http.Post(
		endpoint,
		"application/x-www-form-urlencoded",
		bytes.NewReader(p.toURLEncoded()),
	)

	if err != nil {
		log.Println(err)
	}
}

// Build the struct holding all information about the notification.
func buildHipChat(c *config.Config, j *logging.Job, chl string) hipchatPayload {
	p := hipchatPayload{
		Color:  "green",
		Notify: true,
		Format: "text",
		Room:   chl,
		From:   "Leeroy",
	}

	success := "success"

	if j.Success() == false {
		success = "failed"
		p.Color = "red"
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

	p.Message = m

	return p
}
