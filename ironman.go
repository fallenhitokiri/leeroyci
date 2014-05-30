package main

import (
	"flag"
	"ironman/build"
	"ironman/callbacks"
	"ironman/config"
	"ironman/logging"
	"ironman/web"
	"log"
	"net/http"
)

var cfgFlag = flag.String("config", "ironman.json", "JSON formatted config")

func main() {
	flag.Parse()

	c := config.FromFile(*cfgFlag)

	jobs := make(chan logging.Job, 100)
	b := logging.New(c.BuildLogPath)

	go build.Build(jobs, &c, b)

	log.Println("Ironman up an running!")

	http.HandleFunc("/callback/", func(w http.ResponseWriter, r *http.Request) {
		callbacks.Callback(w, r, jobs, &c, b)
	})
	http.HandleFunc("/status/", func(w http.ResponseWriter, r *http.Request) {
		web.Status(w, r, &c, b)
	})
	http.HandleFunc("/status/repo/", func(w http.ResponseWriter, r *http.Request) {
		web.Repo(w, r, &c, b)
	})
	http.HandleFunc("/status/branch/", func(w http.ResponseWriter, r *http.Request) {
		web.Branch(w, r, &c, b)
	})
	log.Fatal(http.ListenAndServe(":8082", nil))
}
