// Expose the build status over http.
package web

import (
	"html/template"
	"ironman/config"
	"ironman/logging"
	"log"
	"net/http"
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
	log.Println("repo")
}

func Branch(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	log.Println("branch")
}
