// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"leeroy/config"
	"leeroy/logging"
	"log"
)

// Notify send build and deployment notifications for a job.
func Notify(c *config.Config, j *logging.Job, kind string) {
	if kindSupported(kind) == false {
		log.Fatal("unsupported notification type", kind)
		return
	}

	not := notificationFromJob(j, c)
	not.kind = kind

	// always notify the person who comitted
	go email(c, not, j.Email)

	repo, err := c.ConfigForRepo(j.URL)

	if err != nil {
		log.Println("could not find repo", j.URL)
		return
	}

	for _, n := range repo.Notify {
		switch n.Service {
		case "email":
			// Arguments for email are the mail addresses to notify
			for mail := range n.Arguments {
				go email(c, not, mail)
			}
		case "slack":
			go slack(not, n.Arguments["endpoint"], n.Arguments["channel"])
		case "hipchat":
			go hipchat(not, n.Arguments["key"], n.Arguments["channel"])
		case "campfire":
			go campfire(not, n.Arguments["id"], n.Arguments["room"], n.Arguments["key"])
		default:
			log.Println("Notification not supported", n.Service)
		}
	}
}

// Check if kind is a supported notification type.
func kindSupported(kind string) bool {
	for _, k := range kinds {
		if k == kind {
			return true
		}
	}
	return false
}
