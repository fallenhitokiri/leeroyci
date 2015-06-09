package web

import (
	"net/http"
)

func setupGET(w http.ResponseWriter, r *http.Request) {
	tmpl := getTemplates("setup.html")

	tmpl.Execute(w, map[string]string{"Message": "Hello, world!"})
}
