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
	template := "login.html"
	ctx := make(responseContext)

	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			ctx["error"] = err.Error()
			render(w, r, template, ctx)
			return
		}

		decoder := schema.NewDecoder()
		form := new(loginForm)

		err = decoder.Decode(form, r.PostForm)

		if err != nil {
			ctx["error"] = err.Error()
			render(w, r, template, ctx)
			return
		}

		user, err := database.GetUser(form.Email)

		if err != nil {
			ctx["error"] = err.Error()
			render(w, r, template, ctx)
			return
		}

		auth := database.ComparePassword(form.Password, user.Password)

		if auth == false {
			ctx["error"] = "Authentication failed."
			render(w, r, template, ctx)
			return
		}

		sessionID := user.NewSession()

		session, _ := sessionStore.Get(r, "leeroyci")
		session.Values["session_key"] = sessionID
		session.Save(r, w)

		http.Redirect(w, r, "/builds", 302)
		return
	}

	render(w, r, template, ctx)
	return
}
