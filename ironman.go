package main

import (
	"flag"
	"ironman/build"
	"ironman/callbacks"
	"ironman/config"
	"ironman/logging"
	"log"
	"net/http"
)

var cfgFlag = flag.String("config", "ironman.json", "JSON formatted config")

func main() {
	flag.Parse()

	c := config.FromFile(*cfgFlag)

	jobs := make(chan logging.Job, 100)
	b := logging.Buildlog{}

	go build.Build(jobs, &c, &b)

	http.HandleFunc("/callback/", func(w http.ResponseWriter, r *http.Request) {
		callbacks.Callback(w, r, jobs)
	})
	log.Fatal(http.ListenAndServe(":8082", nil))

	log.Println("Ironman up an running!")
}
