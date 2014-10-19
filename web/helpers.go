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
	"strings"
)

// Returns the format for the response. Default is HTML.
func responseFormat(val url.Values) string {
	if val, ok := val["format"]; ok {
		f := strings.Join(val, "")
		return strings.ToLower(f)
	}

	return "html"
}

// Get a template and execute it.
func render(rw http.ResponseWriter, req *http.Request, jobs []logging.Job, c *config.Config, template string) {
	f := responseFormat(req.URL.Query())

	switch f {
	case "json":
		renderJSON(rw, jobs)
	case "html":
		renderHTML(rw, jobs, c, template)
	default:
		log.Println("unsupported render format", f)
	}
}

// Render and write json response.
func renderJSON(rw http.ResponseWriter, jobs []logging.Job) {
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
func renderHTML(rw http.ResponseWriter, jobs []logging.Job, c *config.Config, template string) {
	t, err := templates.Get(template, c)

	if err != nil {
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
