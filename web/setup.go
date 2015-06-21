package web

import (
	"net/http"

	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

type userForm struct {
	Email     string `schema:"email"`
	FirstName string `schema:"first_name"`
	LastName  string `schema:"last_name"`
	Password  string `schema:"password"`

	URL     string `schema:"url"`
	Secret  string `schema:"secret"`
	SSLCert string `schema:"ssl_cert"`
	SSLKey  string `schema:"ssl_key"`

	Host         string `schema:"host"`
	Port         int    `schema:"port"`
	Sender       string `schema:"sender"`
	SMTPUser     string `schema:"smtp_user"`
	SMPTPassword string `schema:"smtp_password"`
}

func viewSetup(w http.ResponseWriter, r *http.Request) {
	template := "setup.html"
	ctx := make(responseContext)

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			ctx["error"] = err.Error()
			render(w, r, template, ctx)
			return
		}

		decoder := schema.NewDecoder()
		form := new(userForm)

		err = decoder.Decode(form, r.PostForm)

		if err != nil {
			ctx["error"] = err.Error()
			render(w, r, template, ctx)
			return
		}

		user, err := database.CreateUser(
			form.Email,
			form.FirstName,
			form.LastName,
			form.Password,
			true,
		)

		if err != nil {
			ctx["error"] = err.Error()
			render(w, r, template, ctx)
			return
		}

		database.AddConfig(
			form.Secret,
			form.URL,
			form.SSLCert,
			form.SSLKey,
		)

		database.AddMailServer(
			form.Host,
			form.Sender,
			form.SMTPUser,
			form.SMPTPassword,
			form.Port,
		)

		if user != nil {
			database.Configured = true
			http.Redirect(w, r, "/login", 302)
		}
	}

	render(w, r, template, ctx)
	return
}
