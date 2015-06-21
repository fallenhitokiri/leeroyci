package web

import (
	"errors"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

type userSettingsForm struct {
	Email       string `schema:"email"`
	FirstName   string `schema:"first_name"`
	LastName    string `schema:"last_name"`
	Password    string `schema:"password"`
	NewPassword string `schema:"new_password"`
}

func (u userSettingsForm) update(request *http.Request) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(userSettingsForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	user := context.Get(request, contextUser).(*database.User)

	auth := database.ComparePassword(form.Password, user.Password)

	if auth == false {
		return errors.New("Username and password do not match.")
	}

	_, err = user.Update(
		form.Email,
		form.FirstName,
		form.LastName,
		form.NewPassword,
		true,
	)

	return err
}

func viewUserSettings(w http.ResponseWriter, r *http.Request) {
	template := "user/settings.html"
	ctx := make(responseContext)

	if r.Method == "POST" {
		err := userSettingsForm{}.update(r)

		if err == nil {
			ctx["message"] = "Update successful."
		} else {
			ctx["error"] = err.Error()
		}
	}

	render(w, r, template, ctx)
}
