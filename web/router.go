package web

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/justinas/alice"
)

var sessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

// Routes returns a new Goriall router.
func Routes() *mux.Router {
	chainUnauth := alice.New(middlewareLogging, middlewarePanic, middlewareNoConfig)
	chainAuth := alice.New(middlewareLogging, middlewarePanic, middlewareNoConfig, middlewareAuth)
	chainAdmin := alice.New(middlewareLogging, middlewarePanic, middlewareNoConfig, middlewareAuth, middlewareAdmin)

	router := mux.NewRouter()
	router.Handle("/setup", chainUnauth.ThenFunc(viewSetup))
	router.Handle("/login", chainUnauth.ThenFunc(viewLogin))
	router.Handle("/logout", chainUnauth.ThenFunc(viewLogout))
	router.Handle("/callback/{service:[a-zA-Z]+}/{secret:[a-zA-Z]+}", chainUnauth.ThenFunc(viewCallback))

	router.Handle("/", chainAuth.ThenFunc(viewListJobs))
	router.Handle("/user/settings", chainAuth.ThenFunc(viewUserSettings))

	router.Handle("/admin/users", chainAdmin.ThenFunc(viewAdminListUsers))
	router.Handle("/admin/user/add", chainAdmin.ThenFunc(viewAdminCreateUser))
	router.Handle("/admin/user/{uid:[0-9]+}", chainAdmin.ThenFunc(viewAdminEditUser))
	router.Handle("/admin/user/delete/{uid:[0-9]+}", chainAdmin.ThenFunc(viewAdminDeleteUser))

	router.Handle("/admin/repositories", chainAdmin.ThenFunc(viewAdminListRepositories))
	router.Handle("/admin/repository/add", chainAdmin.ThenFunc(viewAdminCreateRepository))
	router.Handle("/admin/repository/{rid:[0-9]+}", chainAdmin.ThenFunc(viewAdminEditRepository))
	router.Handle("/admin/repository/delete/{rid:[0-9]+}", chainAdmin.ThenFunc(viewAdminDeleteRepository))

	router.Handle("/admin/repository/{rid:[0-9]+}/notification/add", chainAdmin.ThenFunc(viewAdminCreateNotification))
	router.Handle("/admin/repository/{rid:[0-9]+}/notification/{nid:[0-9]+}", chainAdmin.ThenFunc(viewAdminEditNotification))
	router.Handle("/admin/repository/{rid:[0-9]+}/notification/delete/{nid:[0-9]+}", chainAdmin.ThenFunc(viewAdminDeleteNotification))

	router.Handle("/admin/repository/{rid:[0-9]+}/command/add", chainAdmin.ThenFunc(viewAdminCreateCommand))
	router.Handle("/admin/repository/{rid:[0-9]+}/command/{cid:[0-9]+}", chainAdmin.ThenFunc(viewAdminEditCommand))
	router.Handle("/admin/repository/{rid:[0-9]+}/command/delete/{cid:[0-9]+}", chainAdmin.ThenFunc(viewAdminDeleteCommand))

	// add rice box to serve static files. Do not use the full middleware stack but
	// only the logging handler. We do not want "notConfigured" to run e.x. so we
	// can make the setup look nice.
	box := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	router.Handle("/static/{path:.*}", middlewareLogging(staticFileServer))

	return router
}
