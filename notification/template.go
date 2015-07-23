// Package notification implements sending notifications to users.
package notification

import (
	"bytes"
	"html/template"

	"github.com/GeertJohan/go.rice"

	"github.com/fallenhitokiri/leeroyci/database"
)

// message returns a formatted message to send through a notification system.
// event specifies what happened - tests completed e.x.
// kind specifies the notification system.
func message(job *database.Job, event, kind string) string {
	ctx := map[string]interface{}{
		"TasksFinished":  job.TasksFinished,
		"DeployFinished": job.DeployFinished,
		"Repository":     job.Repository,
		"Branch":         job.Branch,
		"Commit":         job.Commit,
		"CommitURL":      job.CommitURL,
		"Name":           job.Name,
		"Email":          job.Email,
		"CommandLogs":    job.CommandLogs,
	}

	tmpl, err := getTemplate(event, kind)

	if err != nil {
		return ""
	}

	var buffer bytes.Buffer
	tmpl.Execute(&buffer, ctx)
	return buffer.String()
}

// getTemplate returns the template to use for a notification.
func getTemplate(event, kind string) (*template.Template, error) {
	box, err := rice.FindBox("templates")
	if err != nil {
		return nil, err
	}

	name := kind + "-" + event + ".tmpl"

	tmplStr, err := box.String(name)

	if err != nil {
		return nil, err
	}

	tmpl, err := template.New(name).Parse(tmplStr)

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
