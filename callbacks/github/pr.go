// Structs and methods used to process a pull request.
package github

import (
	"encoding/json"
	"log"
	"net/http"
)

// TODO: parse full body, not just the fields needed

type PRCallback struct {
	Number int
	Action string
	PR     PRPullRequest `json:"pull_request"`
}

type PRPullRequest struct {
	Url          string
	Comments_url string
}

// Handle GitHub pull requests.
func handlePR(req *http.Request) {
	b := parseBody(req)

	var pc PRCallback

	err := json.Unmarshal(b, &pc)
	if err != nil {
		log.Println(string(b))
		panic("Could not unmarshal request")
	}

	log.Println(pc)
}
