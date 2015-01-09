// Provide default templates for all notifications.
package notification

var templates = map[string]string{
	"build": build,
}

var build = "Repository: {{.Repo}} Branch: {{.Branch}} by {{.Name}} <{{.Email}}> -> Build {{if .Status}}success{{else}}failed{{end}}\nDetails: {{.Url}}"
