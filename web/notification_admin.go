package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

type notificationAdminForm struct {
	ID        int64  `schema:"ID"`
	Service   string `schema:"service"`
	Arguments string `schema:"arguments"`
}

func (n notificationAdminForm) add(request *http.Request, repo *database.Repository) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(notificationAdminForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	_, err = database.CreateNotification(form.Service, form.Arguments, repo)

	return err
}

func (n notificationAdminForm) update(request *http.Request, not *database.Notification) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(notificationAdminForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	err = not.Update(form.Service, form.Arguments)

	return err
}

func viewAdminCreateNotification(w http.ResponseWriter, r *http.Request) {
	template := "notification/admin/add.html"
	ctx := make(responseContext)

	vars := mux.Vars(r)
	rid := vars["rid"]

	repo, err := database.GetRepositoryByID(rid)

	if err != nil {
		render(w, r, "403.html", make(responseContext)) // TODO: create 500
		return
	}

	ctx["repository"] = repo
	ctx["services"] = []string{
		database.NotificationServiceEmail,
		database.NotificationServiceSlack,
	}

	if r.Method == "POST" {
		err := notificationAdminForm{}.add(r, repo)

		if err != nil {
			ctx["error"] = err.Error()
		} else {
			uri := "/admin/repository/" + rid
			http.Redirect(w, r, uri, 302)
			return
		}
	}

	render(w, r, template, ctx)
}

func viewAdminEditNotification(w http.ResponseWriter, r *http.Request) {
	template := "notification/admin/edit.html"
	ctx := make(responseContext)

	vars := mux.Vars(r)
	rid := vars["rid"]
	nid := vars["nid"]

	repo, err := database.GetRepositoryByID(rid)

	if err != nil {
		render(w, r, "403.html", make(responseContext)) // TODO: create 500
		return
	}

	not, err := database.GetNotification(nid)

	if err != nil {
		render(w, r, "403.html", make(responseContext)) // TODO: create 500
		return
	}

	ctx["repository"] = repo
	ctx["notification"] = not
	ctx["services"] = []string{
		database.NotificationServiceEmail,
		database.NotificationServiceSlack,
	}

	if r.Method == "POST" {
		err := notificationAdminForm{}.update(r, not)

		if err == nil {
			ctx["message"] = "Update successful."
		} else {
			ctx["error"] = err.Error()
		}
	}

	render(w, r, template, ctx)
}

func viewAdminDeleteNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rid := vars["rid"]
	nid := vars["nid"]

	not, err := database.GetNotification(nid)

	if err == nil {
		not.Delete()
	}

	uri := "/admin/repository/" + rid
	http.Redirect(w, r, uri, 302)
}
