package web

import (
	"net/http"

	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

type configAdminForm struct {
	Secret string
	URL    string
	Cert   string
	Key    string
}

func (c configAdminForm) update(request *http.Request) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(configAdminForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	database.UpdateConfig(form.Secret, form.URL, form.Cert, form.Key)
	return nil
}

func viewAdminUpdateConfig(w http.ResponseWriter, r *http.Request) {
	template := "config/admin/update.html"
	ctx := make(responseContext)

	if r.Method == "POST" {
		err := configAdminForm{}.update(r)

		if err == nil {
			ctx["message"] = "Update successful."
		} else {
			ctx["error"] = err.Error()
		}
	}

	config := database.GetConfig()
	ctx["config"] = config

	render(w, r, template, ctx)
}
