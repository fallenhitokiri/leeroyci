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

func viewSetup(w http.ResponseWriter, r *http.Request) (tmpl string, ctx context) {
	tmpl = "setup.html"
	ctx = NewContext(r)

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			ctx["error"] = err.Error()
			return
		}

		decoder := schema.NewDecoder()
		form := new(userForm)

		err = decoder.Decode(form, r.PostForm)

		if err != nil {
			ctx["error"] = err.Error()
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
			http.Redirect(w, r, "/static/leeroy.jpg", 302)
		}
	}

	return
}
