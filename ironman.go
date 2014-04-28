package main

import (
	"ironman/build"
	"ironman/callbacks"
	"ironman/config"
	"log"
	"net/http"
)

func main() {
	not := make(chan callbacks.Notification, 100)
	c := config.Config{}
	b := build.Buildlog{}

	go build.Build(not, &c, &b)

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		callbacks.Callback(w, r, not)
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}
