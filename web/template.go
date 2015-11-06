package web

import (
	"html/template"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/context"
)

type responseContext map[string]interface{}

func render(w http.ResponseWriter, r *http.Request, template string, ctx responseContext) {
	tmpl := getTemplates(template)

	ctx["user"] = context.Get(r, contextUser)

	tmpl.Execute(w, ctx)
}

// getTemplates returns the template 'name' fully prepared for rendering.
func getTemplates(name string) *template.Template {
	box, err := rice.FindBox("templates")
	if err != nil {
		panic(err)
	}

	base, err := box.String("base.html")

	if err != nil {
		panic(err)
	}

	tmpl, err := template.New(name).Parse(base)

	base, err = box.String(name)

	if err != nil {
		panic(err)
	}

	tmpl, err = tmpl.Parse(base)

	if err != nil {
		panic(err)
	}

	return tmpl
}
