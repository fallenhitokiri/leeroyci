package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/fallenhitokiri/leeroyci/github"
	"github.com/fallenhitokiri/leeroyci/gogs"
)

func viewCallback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]

	switch service {
	case "github":
		github.Handle(r)
	case "gogs":
		gogs.Handle(r)
	default:
		log.Println("Service not supported.")
	}
}
