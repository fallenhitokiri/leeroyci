package web

import (
	"errors"
	"net/http"

	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

type loginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (l loginForm) authenticate(request *http.Request) (string, error) {
	err := request.ParseForm()

	if err != nil {
		return "", err
	}

	decoder := schema.NewDecoder()
	form := new(loginForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return "", err
	}

	user, err := database.GetUser(form.Email)

	if err != nil {
		return "", err
	}

	auth := database.ComparePassword(form.Password, user.Password)

	if auth == false {
		return "", errors.New("Username and password do not match.")
	}

	return user.NewSession(), nil
}

func viewLogin(w http.ResponseWriter, r *http.Request) {
	template := "login.html"
	ctx := make(responseContext)

	if r.Method == "POST" {
		sessionID, err := loginForm{}.authenticate(r)

		if err == nil {
			session, _ := sessionStore.Get(r, "leeroyci")
			session.Values["session_key"] = sessionID
			session.Save(r, w)

			http.Redirect(w, r, "/builds", 302)
			return
		}

		ctx["error"] = err.Error()
	}

	render(w, r, template, ctx)
}
