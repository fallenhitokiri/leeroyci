// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

import (
	"bytes"
	"leeroy/config"
	"leeroy/logging"
	"log"
	"text/template"
)

// Define default template names that are used for various notifications.
const (
	KindBuild       = "build"
	KindDeployStart = "deploy-start"
	KindDeployDone  = "deploy-done"
)

// List of all valid notification types.
var kinds = [...]string{KindBuild, KindDeployStart, KindDeployDone}

// Everything related to a notification.
type notification struct {
	Repo    string
	Branch  string
	Name    string
	Email   string
	Status  bool
	URL     string
	kind    string
	message string
}

// Create a notification from a job.
func notificationFromJob(j *logging.Job, c *config.Config) *notification {
	return &notification{
		Repo:   j.URL,
		Branch: j.Branch,
		Name:   j.Name,
		Email:  j.Email,
		Status: j.Success(),
		URL:    j.StatusURL(c.URL),
	}
}

// Render a notification.
func (n *notification) render() {
	t := template.New(n.kind)
	t, err := t.Parse(templates[n.kind])

	if err != nil {
		log.Fatal(err)
	}

	var r bytes.Buffer
	err = t.Execute(&r, n)

	if err != nil {
		log.Fatal(err)
	}

	n.message = r.String()
}
