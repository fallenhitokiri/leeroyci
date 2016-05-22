package gogs

import (
	"log"
	"net/http"
)

// Handle takes the Gogs event from the header and triggeres the matching
// function.
func Handle(req *http.Request) {
	event := req.Header["X-Gogs-Event"][0]

	switch event {
	case "push":
		handlePush(req)
	default:
		log.Println(req.Header)
	}
}
