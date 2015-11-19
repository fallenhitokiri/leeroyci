package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fallenhitokiri/leeroyci/database"
	"github.com/fallenhitokiri/leeroyci/runner"
	"github.com/fallenhitokiri/leeroyci/web"
	"github.com/fallenhitokiri/leeroyci/websocket"
)

func main() {
	database.NewDatabase("", "")
	websocket.NewServer()
	go runner.Runner()

	router := web.Routes()
	config := database.GetConfig()
	
	httpd := &http.Server{
		Addr: port(),
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if config.Cert != "" {
		log.Fatalln(httpd.ListenAndServeTLS(config.Cert, config.Key))
	} else {
		log.Fatalln(httpd.ListenAndServe())
	}
}

func port() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8082"
	}

	return ":" + port
}
