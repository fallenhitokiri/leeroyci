package main

import (
	"log"
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"
	"github.com/fallenhitokiri/leeroyci/runner"
	"github.com/fallenhitokiri/leeroyci/web"
)

func main() {
	database.NewDatabase("", "")
	go runner.Runner()

	router := web.Routes()
	config := database.GetConfig()
	if config.Cert != "" {
		log.Fatalln(http.ListenAndServeTLS(":8082", config.Cert, config.Key, router))
	} else {
		log.Fatalln(http.ListenAndServe(":8082", router))
	}
}
