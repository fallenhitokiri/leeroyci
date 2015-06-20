package web

import (
	"net/http"
)

func viewListAll(w http.ResponseWriter, r *http.Request) (tmpl string, ctx responseContext) {
	tmpl = "builds/list_all.html"
	ctx = NewContext(r)

	return
}
