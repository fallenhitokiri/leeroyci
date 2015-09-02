package main

import (
	"log"
	"net/http"
	"os"

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
		log.Fatalln(http.ListenAndServeTLS(port(), config.Cert, config.Key, router))
	} else {
		log.Fatalln(http.ListenAndServe(port(), router))
	}
}

func port() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8082"
	}

	return ":" + port
}
