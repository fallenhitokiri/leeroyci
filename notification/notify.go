package notification

import (
	"log"

	"github.com/fallenhitokiri/leeroyci/database"
)

// Notify sends all relevant notifications for a job that are configured for
// the jobs repository.
func Notify(job *database.Job, event string) {
	repo, err := database.GetRepositoryByID(job.RepositoryID)

	if err != nil {
		log.Println(err)
		return
	}

	for _, notificaiton := range repo.Notifications {
		switch notificaiton.Service {
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
