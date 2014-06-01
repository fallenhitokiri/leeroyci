// Auth handles the authentication for web requests and callbacks.
package web

import (
	"log"
	"net/http"
	"strings"
)

// Authenticate a request. Currently only URL based authentication is supported.
func Auth(fn func(http.ResponseWriter, *http.Request), s string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		split := strings.Split(r.URL.Path, "/")
		l := len(split)

		// if the URL does not end on a slash we need to subtract 1 form l
		if split[l-1] == "" {
			l = l - 1
		}

		if split[l-1] != s {
			log.Println("wrong secret", r.Host, split[l-1])
			http.Error(w, "wrong secret", 401)
			return
		}

		fn(w, r)
	}
}
