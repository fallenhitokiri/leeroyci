package web

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/fallenhitokiri/leeroyci/database"
)

const limit = 20

// viewListJobs shows a paginated list of all jobs.
func viewListJobs(w http.ResponseWriter, r *http.Request) {
	ctx := make(responseContext)

	offset := 0
	paramOffset := r.URL.Query().Get("offset")

	if len(paramOffset) > 0 {
		offset = stringToInt(paramOffset)
	}

	ctx["jobs"] = database.GetJobs(offset, limit)

	prev, next, first := previous_next_number(offset)

	ctx["previous_offset"] = prev
	ctx["next_offset"] = next
	ctx["first_page"] = first

	render(w, r, "job/list.html", ctx)
}

// viewJobDetail shows a specific job with all related information.
func viewShowJob(w http.ResponseWriter, r *http.Request) {
	template := "job/detail.html"
	ctx := make(responseContext)

	vars := mux.Vars(r)
	jobID, _ := strconv.Atoi(vars["jid"])

	job := database.GetJob(int64(jobID))
	ctx["job"] = job

	render(w, r, template, ctx)
}

// returns the offset for the previous and next page.
func previous_next_number(offset int) (int, int, bool) {
	count := database.NumberOfJobs()
	prev := offset - limit
	next := offset + limit

	if prev < 0 {
		prev = 0
	}

	if next >= count {
		next = -1
	}

	first := false

	if count > limit && next != limit {
		first = true
	}

	return prev, next, first
}
