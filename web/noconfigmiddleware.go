package web

import (
	"net/http"
	"strings"

	"leeroy/database"
)

// notConfigured redirects to /setup if there is no valid configuration
// and the path is not /setup or /static
func notConfiguredHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.String()

		if database.Configured || path == "/setup" || strings.HasPrefix(path, "/static") {
			next.ServeHTTP(w, r)
		} else {
			//c.Redirect(http.StatusFound, "/setup")
			http.Redirect(w, r, "/setup", 302)
			return
		}
	}

	return http.HandlerFunc(fn)
}
