package web

import (
	"net/http"

	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

type mailserverAdminForm struct {
	Host     string
	Sender   string
	Port     int
	User     string
	Password string
}

func (m mailserverAdminForm) update(request *http.Request) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(mailserverAdminForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	database.UpdateMailServer(form.Host, form.Sender, form.User, form.Password, form.Port)
	return nil
}

func viewAdminUpdateMailserver(w http.ResponseWriter, r *http.Request) {
	template := "mailserver/admin/update.html"
	ctx := make(responseContext)

	if r.Method == "POST" {
		err := mailserverAdminForm{}.update(r)

		if err == nil {
			ctx["message"] = "Update successful."
		} else {
			ctx["error"] = err.Error()
		}
	}

	server := database.GetMailServer()
	ctx["server"] = server

	render(w, r, template, ctx)
}
