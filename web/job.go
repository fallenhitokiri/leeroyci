package web

import (
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"
)

const limit = 2

// viewListJobs shows a paginated list of all jobs.
func viewListJobs(w http.ResponseWriter, r *http.Request) {
	ctx := make(responseContext)

	offset := 0
	param := r.URL.Query().Get("offset")

	if len(param) > 0 {
		offset = stringToInt(param)
	}

	ctx["jobs"] = database.GetJobs(offset, limit)

	prev, next := previous_next_number(offset)

	ctx["previous_offset"] = prev
	ctx["next_offset"] = next
	ctx["first_page"] = next != limit

	render(w, r, "job/list_all.html", ctx)
}

// returns the offset for the previous and next page.
func previous_next_number(offset int) (int, int) {
	count := database.NumberOfJobs()
	prev := offset - limit
	next := offset + limit

	if prev < 0 {
		prev = 0
	}

	if next >= count {
		next = -1
	}

	return prev, next
}
