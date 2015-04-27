package main

import (
	"flag"
	"leeroy/build"
	"leeroy/config"
	"leeroy/integrations"
	"leeroy/logging"
	"leeroy/web"
	"log"
	"net/http"
)

var cfgFlag = flag.String("config", "leeroy.json", "JSON formatted config")
var createUser = flag.Bool("createUser", false, "create a new user")
var updateUser = flag.Bool("updateUser", false, "update user information")
var deleteUser = flag.Bool("deleteUser", false, "delete a user")
var listUsers = flag.Bool("listUsers", false, "list all users")

func main() {
	flag.Parse()

	config.FromFile(*cfgFlag)

	err := config.CONFIG.Validate()

	if err != nil {
		log.Fatal("Configuration error: ", err)
	}

	if *createUser == true {
		config.CONFIG.CreateUserCMD()
		return
	}

	if *updateUser == true {
		config.CONFIG.UpdateUserCMD()
		return
	}

	if *deleteUser == true {
		config.CONFIG.DeleteUserCMD()
		return
	}

	if *listUsers == true {
		config.CONFIG.ListUserCMD()
		return
	}

	jobs := make(chan logging.Job, 100)
	logging.New(config.CONFIG.BuildLogPath)

	go build.Build(jobs)

	log.Println("leeroy up an running!")

	http.HandleFunc("/callback/", web.Auth(func(w http.ResponseWriter, r *http.Request) {
		integrations.Callback(w, r, jobs)
	}, config.CONFIG.Secret))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		web.Status(w, r)
	})
	http.HandleFunc("/status/repo/", func(w http.ResponseWriter, r *http.Request) {
		web.Repo(w, r)
	})
	http.HandleFunc("/status/branch/", func(w http.ResponseWriter, r *http.Request) {
		web.Branch(w, r)
	})
	http.HandleFunc("/status/commit/", func(w http.ResponseWriter, r *http.Request) {
		web.Commit(w, r)
	})
	http.HandleFunc("/status/badge/", func(w http.ResponseWriter, r *http.Request) {
		web.Badge(w, r)
	})

	if config.CONFIG.Scheme() == "https" {
		log.Println("HTTPS:", config.CONFIG.URL)
		log.Fatal(http.ListenAndServeTLS(
			config.CONFIG.Host(),
			config.CONFIG.Cert,
			config.CONFIG.Key,
			nil,
		))
	} else {
		log.Println("HTTP:", config.CONFIG.URL)
		log.Fatal(http.ListenAndServe(config.CONFIG.Host(), nil))
	}
}
