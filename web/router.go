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
	chainUnauth := alice.New(middlewareLogging, middlewareNoConfig)
	chainAuth := alice.New(middlewareLogging, middlewareNoConfig, middlewareAuth)
	chainAdmin := alice.New(middlewareLogging, middlewareNoConfig, middlewareAuth, middlewareAdmin)

	router := mux.NewRouter()
	router.Handle("/setup", chainUnauth.ThenFunc(viewSetup))
	router.Handle("/login", chainUnauth.ThenFunc(viewLogin))
	router.Handle("/logout", chainUnauth.ThenFunc(viewLogout))

	router.Handle("/", chainAuth.ThenFunc(viewListAll))
	router.Handle("/user/settings", chainAuth.ThenFunc(viewUserSettings))

	router.Handle("/admin/users", chainAdmin.ThenFunc(viewAdminListUsers))
	router.Handle("/admin/user/add", chainAdmin.ThenFunc(viewAdminCreateUser))
	router.Handle("/admin/user/{uid:[0-9]+}", chainAdmin.ThenFunc(viewAdminEditUser))
	router.Handle("/admin/user/delete/{uid:[0-9]+}", chainAdmin.ThenFunc(viewAdminDeleteUser))

	// add rice box to serve static files. Do not use the full middleware stack but
	// only the logging handler. We do not want "notConfigured" to run e.x. so we
	// can make the setup look nice.
	box := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	router.Handle("/static/{path:.*}", middlewareLogging(staticFileServer))

	return router
}
