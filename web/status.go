// Web exposes build logs through a web interface.
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

	j, n, p := paginatedJobs(blog.Jobs, getParameter(req, "start", "0"))

	r := newResponse(rw, req)
	r.Context.Jobs = j
	r.Context.Next = n
	r.Template = "status"
	r.TemplatePath = c.Templates

	if p != "" {
		r.Context.Previous = p
	}

	r.render()
}

// View to show builds for a specific repository.
func Repo(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	re := splitFirst(req.URL.Path)

	j, n, p := paginatedJobs(blog.JobsForRepo(re), getParameter(req, "start", "0"))

	r := newResponse(rw, req)
	r.Context.Jobs = j
	r.Context.Next = n
	r.Template = "repo"
	r.TemplatePath = c.Templates

	if p != "" {
		r.Context.Previous = p
	}

	r.render()
}

// View to show builds for a specific repository and branch.
func Branch(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	re := splitFirst(req.URL.Path)
	b := splitSecond(req.URL.Path)

	j, n, p := paginatedJobs(blog.JobsForRepoBranch(re, b), getParameter(req, "start", "0"))

	r := newResponse(rw, req)
	r.Context.Jobs = j
	r.Context.Next = n
	r.Template = "branch"
	r.TemplatePath = c.Templates

	if p != "" {
		r.Context.Previous = p
	}

	r.render()
}

// View to show the build for a commit in a repository.
func Commit(rw http.ResponseWriter, req *http.Request, c *config.Config,
	blog *logging.Buildlog) {
	re := splitFirst(req.URL.Path)
	co := splitSecond(req.URL.Path)

	j := blog.JobByCommit(re, co)

	r := newResponse(rw, req)
	r.Context.Jobs = []*logging.Job{j}
	r.Template = "commit"
	r.TemplatePath = c.Templates

	r.render()
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
}
