package web

import (
	"log"
	"net/http"
)

func viewListAll(w http.ResponseWriter, r *http.Request) (tmpl string, ctx context) {
	tmpl = "builds/list_all.html"
	ctx = NewContext(r)

	return
}
