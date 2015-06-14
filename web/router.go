package web

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Routes returns a new Goriall router.
func Routes() *mux.Router {
	mid := alice.New(loggingHandler, notConfiguredHandler)

	router := mux.NewRouter()
	router.Handle("/setup", mid.ThenFunc(setupGET))

	// add rice box to serve static files. Do not use the full middleware stack but
	// only the logging handler. We do not want "notConfigured" to run e.x. so we
	// can make the setup look nice.
	box := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	router.Handle("/static/{path:.*}", loggingHandler(staticFileServer))

	return router
}
