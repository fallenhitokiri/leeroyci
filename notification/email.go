// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"fmt"
	"log"
	"net/mail"

	"github.com/jpoehls/gophermail"

	"github.com/fallenhitokiri/leeroyci/database"
)

// sendEmail sends an email notification. Event specifies which notification to
// send. Valid choices are EVENT_ (see template.go).
func sendEmail(job *database.Job, event string) {
	mailServer := database.GetMailServer()

	htmlMessage := message(job, database.NotificationServiceEmail, event, TypeHTML)
	txtMessage := message(job, database.NotificationServiceEmail, event, TypeText)
	subject := emailSubject(job, event)
	recipient := mail.Address{
		Name:    job.Name,
		Address: job.Email,
	}

	message := gophermail.Message{
		From:     mailServer.From(),
		To:       []mail.Address{recipient},
		Subject:  subject,
		Body:     txtMessage,
		HTMLBody: htmlMessage,
	}

	err := gophermail.SendMail(mailServer.Server(), mailServer.Auth(), &message)

	if err != nil {
		log.Println(err)
	}
}

// emailSubject returns the subject for an email.
func emailSubject(job *database.Job, event string) string {
	if event == EventBuild {
		return fmt.Sprintf("%s/%s build %s", job.Repository.Name, job.Branch, job.Status())
	}

	if event == EventTest {
		return fmt.Sprintf("%s/%s tests %s", job.Repository.Name, job.Branch, job.Status())
	}

	if event == EventDeployStart {
		return fmt.Sprintf("%s/%s deployment started", job.Repository.Name, job.Branch)
	}

	if event == EventDeployEnd {
		return fmt.Sprintf("%s/%s deploy %s", job.Repository.Name, job.Branch, job.Status())
	}

	return "LeeroyCI is confused - not sure which message this is."
}
