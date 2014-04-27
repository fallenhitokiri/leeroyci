package main

import (
	"ironman/build"
	"ironman/callbacks"
	"log"
	"net/http"
)

func main() {
	not := make(chan callbacks.Notification, 100)

	go build.Build(not)

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		callbacks.Callback(w, r, not)
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}
