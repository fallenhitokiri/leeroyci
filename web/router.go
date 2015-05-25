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
	router.Handle("/static/", http.FileServer(rice.MustFindBox("../static").HTTPBox()))

	return router
}
