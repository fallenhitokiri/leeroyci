package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/fallenhitokiri/leeroyci/github"
)

func viewCallback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]

	switch service {
	case "github":
		github.Handle(r)
	default:
		log.Println("Service not supported.")
	}
}
