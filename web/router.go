package web

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/justinas/alice"
)

var sessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

// Routes returns a new Goriall router.
func Routes() *mux.Router {
	chainUnauth := alice.New(middlewareLogging, middlewareNoConfig)
	chainAuth := alice.New(middlewareLogging, middlewareNoConfig, middlewareAuth)

	router := mux.NewRouter()
	router.Handle("/setup", chainUnauth.ThenFunc(viewSetup))
	router.Handle("/login", chainUnauth.ThenFunc(viewLogin))
	router.Handle("/logout", chainUnauth.ThenFunc(viewLogout))
	router.Handle("/builds", chainAuth.ThenFunc(viewListAll))

	// add rice box to serve static files. Do not use the full middleware stack but
	// only the logging handler. We do not want "notConfigured" to run e.x. so we
	// can make the setup look nice.
	box := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	router.Handle("/static/{path:.*}", middlewareLogging(staticFileServer))

	return router
}
