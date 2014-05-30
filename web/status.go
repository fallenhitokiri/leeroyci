// Expose the build status over http.
package web

import (
	"encoding/hex"
	"html/template"
	"ironman/config"
	"ironman/logging"
	"log"
	"net/http"
	"strings"
)

func Status(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	blog.Sort()
	t := template.New("status")
	t, _ = t.Parse(templateStatus)
	t.Execute(
		rw,
		map[string]interface{}{
			"Jobs": blog.Jobs,
		},
	)
}

func Repo(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitRepo(req.URL.Path)

	j := blog.JobsForRepo(r)

	t := template.New("status")
	t, _ = t.Parse(templateStatus)
	t.Execute(
		rw,
		map[string]interface{}{
			"Jobs": j,
		},
	)
}

func Branch(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	log.Println("branch")
}

// Splits a request path and returns the repo name.
func splitRepo(path string) string {
	s := strings.Split(path, "/")[3]
	r, err := hex.DecodeString(s)

	if err != nil {
		log.Println(err)
	}

	return string(r)
}
