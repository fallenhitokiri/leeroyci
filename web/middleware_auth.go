package web

import (
	"net/http"
	"strings"

	"github.com/fallenhitokiri/leeroyci/database"
)

// middlewareAuth tries to get a session key and authenticate a user adding
// the user instance to the context.
func middlewareAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		// path := r.URL.String()

		// if database.Configured || path == "/setup" || strings.HasPrefix(path, "/static") {
		// 	next.ServeHTTP(w, r)
		// } else {
		// 	http.Redirect(w, r, "/setup", 302)
		// 	return
		// }
	}

	return http.HandlerFunc(fn)
}
