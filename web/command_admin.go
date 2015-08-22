package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

type commandAdminForm struct {
	ID      int64  `schema:"ID"`
	Name    string `schema:"name"`
	Kind    string `schema:"kind"`
	Branch  string `schema:"branch"`
	Execute string `schema:"execute"`
}

func (c commandAdminForm) add(request *http.Request, repo *database.Repository) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(commandAdminForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	_, err = database.CreateCommand(
		repo,
		form.Name,
		form.Execute,
		form.Branch,
		form.Kind,
	)

	return err
}

func (c commandAdminForm) update(request *http.Request, com *database.Command) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(commandAdminForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	err = com.Update(form.Name, form.Kind, form.Branch, form.Execute)

	return err
}

func viewAdminCreateCommand(w http.ResponseWriter, r *http.Request) {
	template := "command/admin/add.html"
	ctx := make(responseContext)

	vars := mux.Vars(r)
	rid, _ := strconv.Atoi(vars["rid"])

	repo, err := database.GetRepositoryByID(int64(rid))

	if err != nil {
		render(w, r, "403.html", make(responseContext)) // TODO: create 500
		return
	}

	ctx["repository"] = repo
	ctx["kinds"] = []string{
		database.CommandKindBuild,
		database.CommandKindDeploy,
		database.CommandKindTest,
	}

	if r.Method == "POST" {
		err := commandAdminForm{}.add(r, repo)

		if err != nil {
			ctx["error"] = err.Error()
		} else {
			uri := fmt.Sprintf("/admin/repository/%d", rid)
			http.Redirect(w, r, uri, 302)
			return
		}
	}

	render(w, r, template, ctx)
}

func viewAdminEditCommand(w http.ResponseWriter, r *http.Request) {
	template := "command/admin/edit.html"
	ctx := make(responseContext)

	vars := mux.Vars(r)
	rid, _ := strconv.Atoi(vars["rid"])
	cid := vars["cid"]

	repo, err := database.GetRepositoryByID(int64(rid))

	if err != nil {
		render(w, r, "403.html", make(responseContext)) // TODO: create 500
		return
	}

	com, err := database.GetCommand(cid)

	if err != nil {
		render(w, r, "403.html", make(responseContext)) // TODO: create 500
		return
	}

	ctx["repository"] = repo
	ctx["command"] = com
	ctx["kinds"] = []string{
		database.CommandKindBuild,
		database.CommandKindDeploy,
		database.CommandKindTest,
	}

	if r.Method == "POST" {
		err := commandAdminForm{}.update(r, com)

		if err == nil {
			ctx["message"] = "Update successful."
		} else {
			ctx["error"] = err.Error()
		}
	}

	render(w, r, template, ctx)
}

func viewAdminDeleteCommand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rid := vars["rid"]
	cid := vars["cid"]

	com, err := database.GetCommand(cid)

	if err == nil {
		com.Delete()
	}

	uri := fmt.Sprintf("/admin/repository/", rid)
	http.Redirect(w, r, uri, 302)
}
