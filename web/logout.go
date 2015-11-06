package web

import (
	"net/http"
)

func viewLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "leeroyci")
	session.Values["session_key"] = nil
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
	return
}
