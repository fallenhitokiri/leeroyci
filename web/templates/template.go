// Template implements helpers to either get an externalnor standard template.
package templates

import (
	"errors"
	"html/template"
	"leeroy/config"
)

// Define standard templates to use if no external source is given.
var templates = map[string]string{
	"status": tmpl_standard,
	"repo":   tmpl_standard,
	"branch": tmpl_standard,
	"commit": tmpl_standard,
}

// Get returns a template either from an external source or the default.
func Get(name string, c *config.Config) (*template.Template, error) {
	if c.Templates == "" {
		if tmpl, exists := templates[name]; exists == true {
			t := template.New(name)
			return t.Parse(tmpl)
		} else {
			return nil, errors.New("Standard template does not exist.")
		}
	}
	return nil, nil
}
