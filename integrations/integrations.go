// integrations handles receiving notifications from repository sources like
// GitHub.
package integrations

import (
	"leeroy/config"
	"leeroy/integrations/github"
	"leeroy/logging"
	"log"
	"net/http"
	"strings"
)

func Callback(rw http.ResponseWriter, req *http.Request, jobs chan logging.Job,
	c *config.Config, blog *logging.Buildlog) {
	s := getService(req)

	switch s {
	case "github":
		github.Parse(jobs, req, blog, c)
	default:
		log.Println("serivce", s, "not supported")
	}
}

// Returns the name of the service and the secret key.
func getService(req *http.Request) string {
	return strings.Split(req.URL.Path, "/")[2]
}
