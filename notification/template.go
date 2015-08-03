// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"bytes"
	"html/template"

	"github.com/GeertJohan/go.rice"

	"github.com/fallenhitokiri/leeroyci/database"
)

const (
	EVENT_TEST         = "test"
	EVENT_BUILD        = "build"
	EVENT_DEPLOY_START = "deploy-start"
	EVENT_DEPLOY_END   = "deploy-end"

	TYPE_HTML = "html"
	TYPE_TEXT = "text"
)

// message returns a formatted message to send through a notification system.
// event specifies what happened - tests completed e.x.
// kind specifies the notification system.
func message(job *database.Job, service, event, kind string) string {
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

	tmpl, err := getTemplate(service, event, kind)

	if err != nil {
		return ""
	}

	var buf bytes.Buffer
	tmpl.Execute(&buf, ctx)
	return buf.String()
}

// getTemplate returns the template to use for a notification.
func getTemplate(service, event, kind string) (*template.Template, error) {
	box, err := rice.FindBox("templates")
	if err != nil {
		return nil, err
	}

	name := service + "-" + kind + "-" + event + ".tmpl"

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
