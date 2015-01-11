// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

// Map of all available templates.
// Key = identifier used by LeeroyCI, value = template string.
var templates = map[string]string{
	KindBuild:       templateBuild,
	KindDeployStart: templateDeployStart,
	KindDeployDone:  templateDeployDone,
}

// Template for build notifications.
var templateBuild = "Repository: {{.Repo}} Branch: {{.Branch}} by {{.Name}} <{{.Email}}> -> Build {{if .Status}}success{{else}}failed{{end}}\nDetails: {{.URL}}"

// Template before a deployment task starts.
var templateDeployStart = "Deploying {{.Branch}} for {{.Repo}}"

// Template after a deployment task finished.
var templateDeployDone = "Deployment {{if .Deploy}}successful{{else}}failed{{end}} {{.Branch}} for {{.Repo}}\nDetails: {{.URL}}"
