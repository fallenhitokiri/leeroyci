// Package github integrates everything necessary to test commits, comment on
// pull requests and close them if the build failed.
package github

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Parse a GitHub request body and add it to the build queue.
func Parse(req *http.Request) {
	e := req.Header["X-Github-Event"][0]

	switch e {
	case "push":
		handlePush(req)
	case "pull_request":
		handlePR(req)
	default:
		log.Println("event not supported", e)
	}
}

// Parse the body of a request.
func parseBody(req *http.Request) []byte {
	b, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println(err)
	}

	return b
}
