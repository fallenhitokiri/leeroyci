// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"log"
	"net/http"

	//"github.com/fallenhitokiri/database"
)

// Handle checks if we are dealing with a pull request or a commit and either
// creates a new job in the queue or a PR watcher.
func Handle(req *http.Request) {
	event := req.Header["X-Github-Event"][0]

	log.Println(event)

	switch event {
	case "push":
		handlePush(req)
	case "pull_request":
		//handlePR(req)
	default:
		log.Println("event not supported", event)
	}
}

func handlePR(req *http.Request) {

}
