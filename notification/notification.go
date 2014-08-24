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
		switch n.Service {
		case "email":
			// Arguments for email are the mail addresses to notify
			for _, mail := range n.Arguments {
				go email(c, j, mail)
			}
		case "slack":
			go slack(c, j) // No arguments for Slack
		case "hipchat":
			go hipchat(c, j) // No arguments for HipChat
		case "campfire":
			go campfire(c, j) // No arguments for Campfire
		default:
			log.Println("Notification not supported", n.Service)
		}
	}
}
