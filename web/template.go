package web

import (
	"html/template"
	"log"

	"github.com/GeertJohan/go.rice"
)

func getTemplates(templates ...string) *template.Template {
	// find a rice.Box
	templateBox, err := rice.FindBox("../templates")
	if err != nil {
		log.Fatal(err)
	}

	// get file contents as string
	templateString, err := templateBox.String("base.html")
	if err != nil {
		log.Fatal(err)
	}

	// parse and execute the template
	tmplMessage, err := template.New("setup").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	return tmplMessage
}
