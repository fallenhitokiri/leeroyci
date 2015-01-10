// Package notification handles all notifications for a job. This includes
// build and deployment notifications.
package notification

// Map of all available templates.
// Key = identifier used by LeeroyCI, value = template string.
var templates = map[string]string{
	"build": build,
}

// Template for build notifications.
var build = "Repository: {{.Repo}} Branch: {{.Branch}} by {{.Name}} <{{.Email}}> -> Build {{if .Status}}success{{else}}failed{{end}}\nDetails: {{.Url}}"
