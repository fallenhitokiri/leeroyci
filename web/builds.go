package web

import (
	"net/http"
)

func viewListAll(w http.ResponseWriter, r *http.Request) {
	ctx := make(responseContext)
	ctx["builds"] = []int{1, 2, 3}
	// ctx["next_builds"] =
	//    ctx["previous_builds"] =
	render(w, r, "builds/list_all.html", ctx)
}
