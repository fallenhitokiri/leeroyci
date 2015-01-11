// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

// Map of all available templates.
// Key = identifier used by LeeroyCI, value = template string.
var templates = map[string]string{
	KindBuild: templateBuild,
}

// Template for build notifications.
var templateBuild = "Repository: {{.Repo}} Branch: {{.Branch}} by {{.Name}} <{{.Email}}> -> Build {{if .Status}}success{{else}}failed{{end}}\nDetails: {{.URL}}"
