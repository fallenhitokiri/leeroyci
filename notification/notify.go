// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"leeroy/config"
	"leeroy/logging"
	"log"
)

// Notify send build and deployment notifications for a job.
func Notify(j *logging.Job, kind string) {
	if kindSupported(kind) == false {
		log.Fatalln("unsupported notification type", kind)
	}

	n := notificationFromJob(j)
	n.kind = kind
	n.render()

	// always notify the person who comitted
	go email(n, j.Email)

	repo, err := config.CONFIG.ConfigForRepo(j.URL)

	if err != nil {
		log.Fatalln("could not find repo", j.URL)
	}

	sendNotifications(n, repo.Notify)
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

// Send all notifications which are configured for a repository.
func sendNotifications(n *notification, nots []config.Notify) {
	for _, not := range nots {
		switch not.Service {
		case "email":
			// Arguments for email are the mail addresses to notify
			for mail := range not.Arguments {
				go email(n, mail)
			}
		case "slack":
			go slack(n, not.Arguments["endpoint"], not.Arguments["channel"])
		case "hipchat":
			go hipchat(n, not.Arguments["key"], not.Arguments["channel"])
		case "campfire":
			go campfire(n, not.Arguments["id"], not.Arguments["room"], not.Arguments["key"])
		default:
			log.Println("Notification not supported", not.Service)
		}
	}
}
