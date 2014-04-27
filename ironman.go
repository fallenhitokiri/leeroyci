package main

import (
	"ironman/callbacks"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", callbacks.Callback)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
