// Provide helpers for requests.
package web

import (
	"encoding/hex"
	"encoding/json"
	"html/template"
	"leeroy/logging"
	"log"
	"net/http"
	"strings"
)

// Returns the format for the response.
func responseFormat(req *http.Request) string {
	if val, ok := req.URL.Query()["format"]; ok {
		return strings.Join(val, "")
	}

	return ""
}

// Get a template and execute it.
func render(rw http.ResponseWriter, req *http.Request, jobs []logging.Job) {
	f := responseFormat(req)

	if f == "json" {
		res, err := json.Marshal(jobs)

		if err != nil {
			log.Println("error marshal", err)
			rw.Header().Set("Content-Type", "application/json")
			rw.Write([]byte(`{"error": "marshal not possible"}`))
			req.Body.Close()
		} else {
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(res)
			req.Body.Close()
		}
		return
	}

	t := template.New("status")
	t, _ = t.Parse(standard)
	t.Execute(
		rw,
		map[string]interface{}{
			"Jobs": jobs,
		},
	)
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
