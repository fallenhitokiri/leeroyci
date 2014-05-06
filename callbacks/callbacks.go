// Callbacks handles receiving notifications from repository sources like
// GitHub.
package callbacks

import (
	"io/ioutil"
	"ironman/logging"
	"log"
	"net/http"
	"strings"
)

func Callback(rw http.ResponseWriter, req *http.Request, jobs chan logging.Job) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic("reading")
	}

	s := service(req)

	switch s {
	case "github":
		parseGitHub(jobs, body)
	default:
		log.Println("serivce", s, "not supported")
	}
}

// Returns the name of the service of the callback.
func service(req *http.Request) string {
	path := req.URL.Path[len("/callback/"):]

	// remove slash at the end of the URL if necessary
	if strings.HasSuffix(path, "/") {
		path = strings.Split(path, "/")[0]
	}

	return path
}
