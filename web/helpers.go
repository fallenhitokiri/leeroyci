// Provide helpers for requests.
package web

import (
	"encoding/hex"
	"encoding/json"
	"leeroy/config"
	"leeroy/logging"
	"leeroy/web/templates"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Returns the format for the response. Default is HTML.
func responseFormat(val string) string {
	if val != "" {
		return strings.ToLower(val)
	}

	return "html"
}

// Get a template and execute it.
func render(rw http.ResponseWriter, req *http.Request, jobs []*logging.Job,
	c *config.Config, template string) {
	params, err := url.ParseQuery(req.URL.RawQuery)

	if err != nil {
		log.Fatalln(err)
	}

	var f string
	var j []*logging.Job
	var n int

	_, e := params["format"]
	if e {
		f = responseFormat(params["format"][0])
	} else {
		f = responseFormat("")
	}

	_, e = params["start"]
	if e {
		j, n = paginatedJobs(jobs, params["start"][0])
	} else {
		j, n = paginatedJobs(jobs, "0")
	}

	switch f {
	case "json":
		renderJSON(rw, j, n)
	case "html":
		renderHTML(rw, j, c, template, n)
	default:
		log.Println("unsupported render format", f)
	}
}

// Render and write json response.
func renderJSON(rw http.ResponseWriter, jobs []*logging.Job, next int) {
	res, err := json.Marshal(jobs)

	if err != nil {
		log.Println("error marshal", err)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write([]byte(`{"error": "marshal not possible"}`))
	} else {
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(res)
	}
	return
}

// Render and write HTML response.
func renderHTML(rw http.ResponseWriter, jobs []*logging.Job, c *config.Config,
	template string, next int) {
	t, err := templates.Get(template, c)

	if err != nil {
		log.Println(err)
		http.Error(rw, "500: Error rendering template.", 500)
	} else {
		t.Execute(
			rw,
			map[string]interface{}{
				"Jobs": jobs,
			},
		)
	}
}

// Splits a request path and returns the first part after the endpoint.
// This is usually the hex string of the repository.
func splitFirst(path string) string {
	s := strings.Split(path, "/")[3]
	r, err := hex.DecodeString(s)

	if err != nil {
		log.Println(err)
	}

	return string(r)
}

// Splits a request path and returns the first part after the endpoint.
// This is likely the branch name or commit sha1.
func splitSecond(path string) string {
	return strings.Split(path, "/")[4]
}

func paginatedJobs(jobs []*logging.Job, start string) ([]*logging.Job, int) {
	c := len(jobs) - 1
	f, err := strconv.Atoi(start)

	if err != nil {
		return jobs, c
	}

	p := 10    // default pagination is 10 jobs
	n := f + p // next start is current start + paginate count

	if f > 0 {
		f -= 1
	}

	l := f + p // default for last slice index is start + paginate count

	// handle start out of range
	if f > c {
		return jobs, 0
	}

	// if the end would be out of range set a new end and return no next
	if l >= c {
		l = f + c - f
		n = 0
	}

	return jobs[f:l], n
}
