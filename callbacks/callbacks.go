// Callbacks handles receiving notifications from repository sources like
// GitHub.
package callbacks

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Notification interface {
	Branch() string
	URL() string
	By() (string, string)
	ShouldBuild() bool
	Commit() string
}

func Callback(rw http.ResponseWriter, req *http.Request, not chan Notification) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		panic("reading")
	}

	s := service(req)

	switch s {
	case "github":
		parseGitHub(not, body)
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
