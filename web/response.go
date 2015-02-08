// Web exposes build logs through a web interface.
package web

import (
	"encoding/json"
	"leeroy/logging"
	"leeroy/web/templates"
	"log"
	"net/http"
	"strings"
)

var (
	formatHTML = "html"
	formatJSON = "json"
)

// Context holds all varibales that could be used by a template.
type Context struct {
	Jobs     []*logging.Job
	Next     string
	Previous string
	URL      string
}

// Response takes care of rendering a template and responding to a request.
type Response struct {
	Writer       http.ResponseWriter
	Template     string
	TemplatePath string
	Format       string
	Context      *Context
}

// newRepsonse sets response format and returns a Response.
func newResponse(rw http.ResponseWriter, req *http.Request) *Response {
	r := &Response{
		Context: &Context{},
		Writer:  rw,
	}

	r.Format = strings.ToLower(getParameter(req, "format", formatHTML))

	return r
}

// render a response.
func (r *Response) render() {
	switch r.Format {
	case formatHTML:
		r.renderHTML()
	case formatJSON:
		r.renderJSON()
	default:
		log.Fatalln("unsupported render format: ", r.Format)
	}
}

// render a response as JSON.
func (r *Response) renderJSON() {
	r.Writer.Header().Set("Content-Type", "application/json")

	if res, err := json.Marshal(r.Context.Jobs); err != nil {
		log.Println("error marshal", err)
		r.Writer.Write([]byte(`{"error": "marshal not possible"}`))
	} else {
		r.Writer.Write(res)
	}
}

// render a response as HTML.
func (r *Response) renderHTML() {
	if t, err := templates.Get(r.Template, r.TemplatePath); err != nil {
		log.Println(err)
		http.Error(r.Writer, "500: Error rendering template.", 500)
	} else {
		t.Execute(r.Writer, r.Context)
	}
}
