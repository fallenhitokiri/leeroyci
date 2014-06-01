// Callbacks handles receiving notifications from repository sources like
// GitHub.
package callbacks

import (
	"leeroy/callbacks/github"
	"leeroy/config"
	"leeroy/logging"
	"log"
	"net/http"
	"strings"
)

func Callback(rw http.ResponseWriter, req *http.Request, jobs chan logging.Job,
	c *config.Config, blog *logging.Buildlog) {
	s, k := splitUrl(req)

	if k != c.Secret {
		log.Println("wrong key from", req.Host)
		return
	}

	switch s {
	case "github":
		github.Parse(jobs, req, blog, c)
	default:
		log.Println("serivce", s, "not supported")
	}
}

// Returns the name of the service and the secret key.
func splitUrl(req *http.Request) (string, string) {
	split := strings.Split(req.URL.Path, "/")
	s := split[2]
	k := split[3]

	return s, k
}
