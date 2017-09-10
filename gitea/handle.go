package gitea

import (
	"log"
	"net/http"
)

func Handle(req *http.Request) {
	event := req.Header.Get("X-Gogs-Event")

	switch event {
	case "push":
		handlePush(req)
	default:
		log.Println("event not supported", event)
	}
}
