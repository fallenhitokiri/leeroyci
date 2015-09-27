package web

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/justinas/alice"

	"github.com/fallenhitokiri/leeroyci/websocket"
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
	router.Handle("/callback/{service:[a-zA-Z]+}/{secret:[a-zA-Z0-9]+}", chainUnauth.ThenFunc(viewCallback))

	router.Handle("/", chainAuth.ThenFunc(viewListJobs)).Name("listJobs")
	router.Handle("/{jid:[0-9]+}", chainAuth.ThenFunc(viewDetailJob))
	router.Handle("/{jid:[0-9]+}/cancel", chainAuth.ThenFunc(viewCancelJob))
	router.Handle("/{jid:[0-9]+}/rerun", chainAuth.ThenFunc(viewRerunJob))
	router.Handle("/search", chainAuth.ThenFunc(viewSearchJobs))

	router.Handle("/user/settings", chainAuth.ThenFunc(viewUpdateUser))

	router.Handle("/admin/users", chainAdmin.ThenFunc(viewAdminListUsers))
	router.Handle("/admin/user/create", chainAdmin.ThenFunc(viewAdminCreateUser))
	router.Handle("/admin/user/{uid:[0-9]+}", chainAdmin.ThenFunc(viewAdminUpdateUser))
	router.Handle("/admin/user/delete/{uid:[0-9]+}", chainAdmin.ThenFunc(viewAdminDeleteUser))

	router.Handle("/admin/repositories", chainAdmin.ThenFunc(viewAdminListRepositories))
	router.Handle("/admin/repository/create", chainAdmin.ThenFunc(viewAdminCreateRepository))
	router.Handle("/admin/repository/{rid:[0-9]+}", chainAdmin.ThenFunc(viewAdminUpdateRepository))
	router.Handle("/admin/repository/delete/{rid:[0-9]+}", chainAdmin.ThenFunc(viewAdminDeleteRepository))

	router.Handle("/admin/mailserver/update", chainAdmin.ThenFunc(viewAdminUpdateMailserver))

	router.Handle("/admin/config/update", chainAdmin.ThenFunc(viewAdminUpdateConfig))

	router.Handle("/admin/repository/{rid:[0-9]+}/notification/create", chainAdmin.ThenFunc(viewAdminCreateNotification))
	router.Handle("/admin/repository/{rid:[0-9]+}/notification/{nid:[0-9]+}", chainAdmin.ThenFunc(viewAdminUpdateNotification))
	router.Handle("/admin/repository/{rid:[0-9]+}/notification/delete/{nid:[0-9]+}", chainAdmin.ThenFunc(viewAdminDeleteNotification))

	router.Handle("/admin/repository/{rid:[0-9]+}/command/create", chainAdmin.ThenFunc(viewAdminCreateCommand))
	router.Handle("/admin/repository/{rid:[0-9]+}/command/{cid:[0-9]+}", chainAdmin.ThenFunc(viewAdminUpdateCommand))
	router.Handle("/admin/repository/{rid:[0-9]+}/command/delete/{cid:[0-9]+}", chainAdmin.ThenFunc(viewAdminDeleteCommand))

	// add rice box to serve static files. Do not use the full middleware stack but
	// only the logging handler. We do not want "notConfigured" to run e.x. so we
	// can make the setup look nice.
	box := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	router.Handle("/static/{path:.*}", middlewareLogging(staticFileServer))

	router.Handle("/websocket", websocket.GetHandler())

	return router
}
