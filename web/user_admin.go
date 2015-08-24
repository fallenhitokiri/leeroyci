package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

// userAdminForm is the form used by admins to edit users.
type userAdminForm struct {
	Email     string `schema:"email"`
	FirstName string `schema:"first_name"`
	LastName  string `schema:"last_name"`
	Password  string `schema:"password"`
	Admin     bool   `schema:"is_admin"`
}

// create creates a new user in the system.
func (u userAdminForm) create(request *http.Request) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(userAdminForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	_, err = database.CreateUser(
		form.Email,
		form.FirstName,
		form.LastName,
		form.Password,
		form.Admin,
	)

	return err
}

// update updates an existing user based on the form.
func (u userAdminForm) update(request *http.Request, uid string) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(userAdminForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	user, err := database.GetUserByID(uid)

	if err != nil {
		return err
	}

	_, err = user.Update(
		form.Email,
		form.FirstName,
		form.LastName,
		form.Password,
		form.Admin,
	)

	return err
}

// viewAdminListUsers lists all users in the system.
func viewAdminListUsers(w http.ResponseWriter, r *http.Request) {
	template := "user/admin/list.html"
	ctx := make(responseContext)

	ctx["users"] = database.ListUsers()

	render(w, r, template, ctx)
}

// viewAdminCreateUser creates a new user.
func viewAdminCreateUser(w http.ResponseWriter, r *http.Request) {
	template := "user/admin/create.html"
	ctx := make(responseContext)

	if r.Method == "POST" {
		err := userAdminForm{}.create(r)

		if err != nil {
			ctx["error"] = err.Error()
		} else {
			http.Redirect(w, r, "/admin/users", 302)
			return
		}
	}

	render(w, r, template, ctx)
}

// viewAdminUpdateUser edits a user for a given uid.
func viewAdminUpdateUser(w http.ResponseWriter, r *http.Request) {
	template := "user/admin/update.html"
	ctx := make(responseContext)

	vars := mux.Vars(r)
	uid := vars["uid"]

	if r.Method == "POST" {
		err := userAdminForm{}.update(r, uid)

		if err == nil {
			ctx["message"] = "Update successful."
		} else {
			ctx["error"] = err.Error()
		}
	}

	user, err := database.GetUserByID(uid)

	if err != nil {
		ctx["error"] = err.Error()
	} else {
		ctx["edit_user"] = user
	}

	render(w, r, template, ctx)
}

// viewAdminDeleteUser deletes a user for a given uid.
func viewAdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	user, err := database.GetUserByID(uid)

	if err == nil {
		user.Delete()
	}

	http.Redirect(w, r, "/admin/users", 302)
}
