package web

import (
	"net/http"

	"github.com/gorilla/context"

	"github.com/fallenhitokiri/leeroyci/database"
)

// middlewareAuth tries to get a session key and authenticate a user adding
// the user instance to the context.
func middlewareAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessionStore.Get(r, "leeroyci")
		sessionKey := session.Values["session_key"]

		if sessionKey == nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		user, err := database.GetUserBySession(sessionKey.(string))

		if err == nil {
			context.Set(r, contextUser, user)
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/login", 302)
			return
		}
	}

	return http.HandlerFunc(fn)
}
