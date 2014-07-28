// Expose the build status over http.
package web

import (
	"encoding/hex"
	"html/template"
	"leeroy/config"
	"leeroy/logging"
	"log"
	"net/http"
	"strings"
)

// View that shows all builds ever done.
func Status(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	blog.Sort()

	render(rw, blog.Jobs)
}

// View to show builds for a specific repository.
func Repo(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitFirst(req.URL.Path)

	j := blog.JobsForRepo(r)

	render(rw, j)
}

// View to show builds for a specific repository and branch.
func Branch(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitFirst(req.URL.Path)
	b := splitSecond(req.URL.Path)

	j := blog.JobsForRepoBranch(r, b)

	render(rw, j)
}

// View to show the build for a commit in a repository.
func Commit(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitFirst(req.URL.Path)
	co := splitSecond(req.URL.Path)

	j := blog.JobByCommit(r, co)

	render(rw, []logging.Job{j})
}

// Endpoint returning a badge showing the build status for a repository and
// branch. It returns an SVG.
func Badge(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitFirst(req.URL.Path)
	b := splitSecond(req.URL.Path)

	j := blog.JobsForRepoBranch(r, b)

	var svg []byte

	if len(j) == 0 {
		svg = []byte(badgeNoResults)
	} else if j[0].Success() {
		svg = []byte(badgeSuccess)
	} else {
		svg = []byte(badgeFailed)
	}

	rw.Header().Set("Content-Type", "image/svg+xml")
	rw.Write(svg)
	req.Body.Close()
}

// Get a template and execute it.
func render(rw http.ResponseWriter, jobs []logging.Job) {
	t := template.New("status")
	t, _ = t.Parse(standard)
	t.Execute(
		rw,
		map[string]interface{}{
			"Jobs": jobs,
		},
	)
}

// Splits a request path and returns the first part after the endpoint.
// This is usually the hex string of the repository.
func splitFirst(path string) string {
	s := strings.Split(path, "/")[3]
	r, err := hex.DecodeString(s)

	if err != nil {
		log.Println(err)
	}

	return string(r)
}

// Splits a request path and returns the first part after the endpoint.
// This is likely the branch name or commit sha1.
func splitSecond(path string) string {
	return strings.Split(path, "/")[4]
}
