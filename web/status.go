// Expose the build status over http.
package web

import (
	"leeroy/config"
	"leeroy/logging"
	"net/http"
)

// View that shows all builds ever done.
func Status(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	blog.Sort()

	render(rw, req, blog.Jobs, c, "status")
}

// View to show builds for a specific repository.
func Repo(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitFirst(req.URL.Path)

	j := blog.JobsForRepo(r)

	render(rw, req, j, c, "repo")
}

// View to show builds for a specific repository and branch.
func Branch(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitFirst(req.URL.Path)
	b := splitSecond(req.URL.Path)

	j := blog.JobsForRepoBranch(r, b)

	render(rw, req, j, c, "branch")
}

// View to show the build for a commit in a repository.
func Commit(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	r := splitFirst(req.URL.Path)
	co := splitSecond(req.URL.Path)

	j := blog.JobByCommit(r, co)

	render(rw, req, []logging.Job{j}, c, "commit")
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
