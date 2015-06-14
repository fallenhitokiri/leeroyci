package web

import (
	"html/template"
	"log"

	"github.com/GeertJohan/go.rice"
)

// getTemplates returns the template 'name' fully prepared for rendering.
func getTemplates(name string) *template.Template {
	box, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}

	base, err := box.String("base.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New(name).Parse(base)

	base, err = box.String(name)

	if err != nil {
		log.Fatal(err)
	}

	tmpl, err = tmpl.Parse(base)

	if err != nil {
		log.Fatal(err)
	}

	return tmpl
}
