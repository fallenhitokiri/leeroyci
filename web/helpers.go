// Provide helpers for requests.
package web

import (
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
