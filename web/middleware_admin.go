package web

import (
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"
)

// middlewareAdmin checks if the authenticated user is an admin. If this is not
// the case we raise an 403
func middlewareAdmin(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessionStore.Get(r, "leeroyci")
		sessionKey := session.Values["session_key"]

		if sessionKey == nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		user, err := database.GetUserBySession(sessionKey.(string))

		if err == nil {
			if user.Admin == false {
				render(w, r, "403.html", make(responseContext))
				return
			}
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/login", 302)
			return
		}
	}

	return http.HandlerFunc(fn)
}
