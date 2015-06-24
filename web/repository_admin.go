package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	"github.com/fallenhitokiri/leeroyci/database"
)

type repositoryAdminForm struct {
	Name string `schema:"name"`
	URL  string `schema:"url"`

	CommentPR bool   `schema:"comment_pr"`
	ClosePR   bool   `schema:"close_pr"`
	StatusPR  bool   `schema:"status_pr"`
	AccessKey string `schema:"access_key"`
}

func (r repositoryAdminForm) add(request *http.Request) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(repositoryAdminForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	_, err = database.CreateRepository(
		form.Name,
		form.URL,
		form.AccessKey,
		form.CommentPR,
		form.ClosePR,
		form.StatusPR,
	)

	return err
}

func (r repositoryAdminForm) update(request *http.Request, rid string) error {
	err := request.ParseForm()

	if err != nil {
		return err
	}

	decoder := schema.NewDecoder()
	form := new(repositoryAdminForm)

	err = decoder.Decode(form, request.PostForm)

	if err != nil {
		return err
	}

	repo, err := database.GetRepositoryByID(rid)

	if err != nil {
		return err
	}

	_, err = repo.Update(
		form.Name,
		form.URL,
		form.AccessKey,
		form.CommentPR,
		form.ClosePR,
		form.StatusPR,
	)

	return err
}

func viewAdminListRepositories(w http.ResponseWriter, r *http.Request) {
	template := "repository/admin/list.html"
	ctx := make(responseContext)

	ctx["repositories"] = database.ListRepositories()

	render(w, r, template, ctx)
}

func viewAdminCreateRepository(w http.ResponseWriter, r *http.Request) {
	template := "repository/admin/add.html"
	ctx := make(responseContext)

	if r.Method == "POST" {
		err := repositoryAdminForm{}.add(r)

		if err != nil {
			ctx["error"] = err.Error()
		} else {
			http.Redirect(w, r, "/admin/repositories", 302)
			return
		}
	}

	render(w, r, template, ctx)
}

func viewAdminEditRepository(w http.ResponseWriter, r *http.Request) {
	template := "repository/admin/edit.html"
	ctx := make(responseContext)

	vars := mux.Vars(r)
	rid := vars["rid"]

	if r.Method == "POST" {
		err := repositoryAdminForm{}.update(r, rid)

		if err == nil {
			ctx["message"] = "Update successful."
		} else {
			ctx["error"] = err.Error()
		}
	}

	repo, err := database.GetRepositoryByID(rid)

	if err != nil {
		ctx["error"] = err.Error()
	} else {
		ctx["repository"] = repo
	}

	render(w, r, template, ctx)
}

func viewAdminDeleteRepository(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rid := vars["rid"]

	repo, err := database.GetRepositoryByID(rid)

	if err == nil {
		repo.Delete()
	}

	http.Redirect(w, r, "/admin/repositories", 302)
}
