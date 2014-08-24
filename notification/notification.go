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
			for _, mail := range n.Arguments {
				go email(c, j, mail)
			}
			continue
		}

		if n.Service == "slack" {
			// No arguments for Slack
			go slack(c, j)
			continue
		}

		if n.Service == "hipchat" {
			// No arguments for HipChat
			go hipchat(c, j)
			continue
		}

		if n.Service == "campfire" {
			// No arguments for Campfire
			go campfire(c, j)
			continue
		}
	}
}
