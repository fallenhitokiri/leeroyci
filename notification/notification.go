// Notifications to send after a build process finished.
package notification

import (
	"leeroy/config"
	"leeroy/logging"
	"log"
)

// Run all notifications configured for a
func Notify(c *config.Config, j *logging.Job) {
	// always notify the person who comitted
	go email(c, j, j.Email)

	repo, err := c.ConfigForRepo(j.URL)

	if err != nil {
		log.Println("could not find repo", j.URL)
		return
	}

	for _, n := range repo.Notify {
		if n.Service == "email" {
			// Arguments for email are the mail addresses to notify
			for mail, _ := range n.Arguments {
				go email(c, j, mail)
			}
			continue
		}

		if n.Service == "slack" {
			go slack(c, j, n.Arguments["endpoint"], n.Arguments["channel"])
			continue
		}

		if n.Service == "hipchat" {
			go hipchat(c, j, n.Arguments["key"], n.Arguments["channel"])
			continue
		}
	}
}
