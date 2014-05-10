// Callbacks handles receiving notifications from repository sources like
// GitHub.
package callbacks

import (
	"io/ioutil"
	"ironman/callbacks/github"
	"ironman/logging"
	"log"
	"net/http"
	"strings"
)

func Callback(rw http.ResponseWriter, req *http.Request, jobs chan logging.Job,
	secret string) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic("reading")
	}

	s, k := splitUrl(req)

	if k != secret {
		log.Println("wrong key from", req.Host)
		return
	}

	switch s {
	case "github":
		github.Parse(jobs, body)
	default:
		log.Println("serivce", s, "not supported")
	}
}

// Returns the name of the service and the secret key.
func splitUrl(req *http.Request) (string, string) {
	path := req.URL.Path[len("/callback/"):]

	k := strings.Split(path, "/")[0]
	s := strings.Split(path, "/")[1]

	return s, k
}
