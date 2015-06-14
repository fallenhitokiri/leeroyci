package web

import (
	"net/http"

	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

type loginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func viewLogin(w http.ResponseWriter, r *http.Request) {
	tmpl := getTemplates("login.html")
	var payload = make(map[string]interface{})

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			payload["error"] = err.Error()
			tmpl.Execute(w, payload)
			return
		}

		decoder := schema.NewDecoder()
		form := new(loginForm)

		err = decoder.Decode(form, r.PostForm)

		if err != nil {
			payload["error"] = err.Error()
			tmpl.Execute(w, payload)
			return
		}

		user, err := database.GetUser(form.Email)

		if err != nil {
			payload["error"] = err.Error()
			tmpl.Execute(w, payload)
			return
		}

		auth := database.ComparePassword(form.Password, user.Password)

		if auth == false {
			payload["error"] = err.Error()
			tmpl.Execute(w, payload)
			return
		}

		session, _ := store.Get(r, "leeroyci")
		session.Values["admin"] = user.Admin
		session.Save(r, w)
	}

	tmpl.Execute(w, payload)
}
