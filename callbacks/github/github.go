// GitHub provides all structs to unmarshal a GitHub webhook.
package github

import (
	"io/ioutil"
	"ironman/logging"
	"log"
	"net/http"
)

// Parse a GitHub request body and add it to the build queue.
func Parse(jobs chan logging.Job, req *http.Request) {
	e := req.Header["X-Github-Event"][0]

	switch e {
	case "push":
		handlePush(req, jobs)
	default:
		log.Println("event not supported", e)
	}
}

// Parse the body of a request.
func parseBody(req *http.Request) []byte {
	b, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic("reading")
	}

	return b
}
