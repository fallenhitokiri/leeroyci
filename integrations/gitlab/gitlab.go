// GitLab integration.
package gitlab

import (
	"encoding/json"
	"io/ioutil"
	"leeroy/config"
	"leeroy/logging"
	"log"
	"net/http"
)

// Helper struct to figure out what kind of push we received. GitLab does
// not send any identifier but a different payload. If kind is empty it is
// a push, if it is not empty it is a merge request or issue.
type Identifier struct {
	Kind string `json:"object_kind"`
}

// Parse a GitLab request body and add it to the build queue.
func Parse(jobs chan logging.Job, req *http.Request, blog *logging.Buildlog,
	c *config.Config) {

	var i Identifier
	pb := parseBody(req)
	json.Unmarshal(pb, &i)

	switch i.Kind {
	case "":
		handlePush(pb, jobs)
	case "merge_request":
		log.Println("marge request")
	default:
		log.Println("event not supported", i.Kind)
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
