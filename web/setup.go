package web

import (
	"net/http"

	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

type configForm struct {
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

func (c configForm) save(request *http.Request) (*database.User, error) {
	err := request.ParseForm()

	if err != nil {
		return nil, err
	}

	decoder := schema.NewDecoder()
	form := new(configForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return nil, err
	}

	user, err := database.CreateUser(
		form.Email,
		form.FirstName,
		form.LastName,
		form.Password,
		true,
	)

	if err != nil {
		return nil, err
	}

	database.AddConfig(
		form.Secret,
		form.URL,
		form.SSLCert,
		form.SSLKey,
		1,
	)

	database.AddMailServer(
		form.Host,
		form.Sender,
		form.SMTPUser,
		form.SMPTPassword,
		form.Port,
	)

	return user, nil
}

func viewSetup(w http.ResponseWriter, r *http.Request) {
	template := "setup.html"
	ctx := make(responseContext)

	if r.Method == "POST" {
		_, err := configForm{}.save(r)

		if err == nil {
			database.Configured = true
			http.Redirect(w, r, "/login", 302)
			return
		}

		ctx["error"] = err.Error()
	}

	render(w, r, template, ctx)
}
