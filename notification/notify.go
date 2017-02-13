package notification

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/fallenhitokiri/leeroyci/database"
)

// Notify sends all relevant notifications for a job that are configured for
// the jobs repository.
func Notify(job *database.Job, event string) {
	log.Println("Start notify")
	log.Println("Job:", spew.Dump(job))
	repo, err := database.GetRepositoryByID(job.RepositoryID)
	log.Println("Name:", repo.Name)

	if err != nil {
		log.Println(err)
		return
	}

	sendWebsocket(job, event)
	for _, notification := range repo.Notifications {
		switch notification.Service {
		case database.NotificationServiceEmail:
			sendEmail(job, event)
		case database.NotificationServiceSlack:
			sendSlack(job, event)
		case database.NotificationServiceHipchat:
			sendHipchat(job, event)
		case database.NotificationServiceCampfire:
			sendCampfire(job, event)
		default:
			continue
		}
	}
}
