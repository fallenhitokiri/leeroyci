// Package web implements the complete web interface for LeeroyCI. This includes
// exposing the build log to the web and implementing different actions like
// rerunning jobs, deploying or administrative tasks.
package web

import (
	"leeroy/database"
	"net/http"
)

// Status shows all builds ever done.
func Status(rw http.ResponseWriter, req *http.Request) {
	jobs := database.GetAllJobs()
	config := database.GetConfig()

	j, n, p := paginatedJobs(jobs, getParameter(req, "start", "0"))

	r := newResponse(rw, req)
	r.Context.Jobs = j
	r.Context.Next = n
	r.Context.URL = config.URL
	r.Template = "status"
	r.TemplatePath = config.Templates

	if p != "" {
		r.Context.Previous = p
	}

	r.render()
}

// Repo shows builds for a specific repository.
func Repo(rw http.ResponseWriter, req *http.Request) {
	//re := splitFirst(req.URL.Path)
	config := database.GetConfig()

	j, n, p := paginatedJobs(
		database.GetAllJobs(), // TODO: query jobs for re
		getParameter(req, "start", "0"),
	)

	r := newResponse(rw, req)
	r.Context.Jobs = j
	r.Context.Next = n
	r.Context.URL = config.URL
	r.Template = "repo"
	r.TemplatePath = config.Templates

	if p != "" {
		r.Context.Previous = p
	}

	r.render()
}

// Branch shows builds for a specific repository and branch.
func Branch(rw http.ResponseWriter, req *http.Request) {
	//re := splitFirst(req.URL.Path)
	//b := splitSecond(req.URL.Path)
	config := database.GetConfig()

	j, n, p := paginatedJobs(
		database.GetAllJobs(), // TODO: query jobs for re + b
		getParameter(req, "start", "0"),
	)

	r := newResponse(rw, req)
	r.Context.Jobs = j
	r.Context.Next = n
	r.Context.URL = config.URL
	r.Template = "branch"
	r.TemplatePath = config.Templates

	if p != "" {
		r.Context.Previous = p
	}

	r.render()
}

// Commit shows the build for a commit in a repository.
func Commit(rw http.ResponseWriter, req *http.Request) {
	//re := splitFirst(req.URL.Path)
	//co := splitSecond(req.URL.Path)
	config := database.GetConfig()

	j := database.GetAllJobs()[0] //logging.BUILDLOG.JobByCommit(re, co)

	r := newResponse(rw, req)
	r.Context.Jobs = []*database.Job{j}
	r.Context.URL = config.URL
	r.Template = "commit"
	r.TemplatePath = config.Templates

	r.render()
}

// Badge returns a badge - as SVG - showing the build status for a repository and
// branch.
func Badge(rw http.ResponseWriter, req *http.Request) {
	//r := splitFirst(req.URL.Path)
	//b := splitSecond(req.URL.Path)

	j := database.GetAllJobs()[0] //logging.BUILDLOG.JobsForRepoBranch(r, b)

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
