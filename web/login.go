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

func viewLogin(w http.ResponseWriter, r *http.Request) (tmpl string, ctx context) {
	tmpl = "login.html"
	ctx = make(context)

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			ctx["error"] = err.Error()
			return
		}

		decoder := schema.NewDecoder()
		form := new(loginForm)

		err = decoder.Decode(form, r.PostForm)

		if err != nil {
			ctx["error"] = err.Error()
			return
		}

		user, err := database.GetUser(form.Email)

		if err != nil {
			ctx["error"] = err.Error()
			return
		}

		auth := database.ComparePassword(form.Password, user.Password)

		if auth == false {
			ctx["error"] = "Authentication failed."
			return
		}

		session, _ := store.Get(r, "leeroyci")
		session.Values["admin"] = user.Admin
		session.Save(r, w)
	}

	return
}
