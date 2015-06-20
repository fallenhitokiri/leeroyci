package web

import (
	"net/http"
	"strings"

	"github.com/fallenhitokiri/leeroyci/database"
)

// middlewareNoConfig redirects to /setup if there is no valid configuration
// and the path is not /setup or /static
func middlewareNoConfig(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.String()

		if database.Configured || path == "/setup" || strings.HasPrefix(path, "/static") {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/setup", 302)
			return
		}
	}

	return http.HandlerFunc(fn)
}
