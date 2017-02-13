package web

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/fallenhitokiri/leeroyci/websocket"
)

var sessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

func chainUnauth(next http.HandlerFunc) http.Handler {
	return middlewareLogging(middlewarePanic(middlewareNoConfig(next)))
}

func chainAuth(next http.HandlerFunc) http.Handler {
	return middlewareLogging(middlewarePanic(middlewareNoConfig(middlewareAuth(next))))
}
func chainAdmin(next http.HandlerFunc) http.Handler {
	return middlewareLogging(middlewarePanic(middlewareNoConfig(middlewareAuth(middlewareAdmin(next)))))
}

// Routes returns a new Goriall router.
func Routes() *mux.Router {

	router := mux.NewRouter()
	router.Handle("/setup", chainUnauth(viewSetup))
	router.Handle("/login", chainUnauth(viewLogin))
	router.Handle("/logout", chainUnauth(viewLogout))
	router.Handle("/callback/{service:[a-zA-Z]+}/{secret:[a-zA-Z0-9-]+}", chainUnauth(viewCallback))

	router.Handle("/", chainAuth(viewListJobs))
	router.Handle("/{jid:[0-9]+}", chainAuth(viewDetailJob))
	router.Handle("/{jid:[0-9]+}/cancel", chainAuth(viewCancelJob))
	router.Handle("/{jid:[0-9]+}/rerun", chainAuth(viewRerunJob))
	router.Handle("/search", chainAuth(viewSearchJobs))

	router.Handle("/user/settings", chainAuth(viewUpdateUser))
	router.Handle("/user/regenerate-accesskey", chainAuth(viewRegenrateAccessKey))

	router.Handle("/admin/users", chainAdmin(viewAdminListUsers))
	router.Handle("/admin/user/create", chainAdmin(viewAdminCreateUser))
	router.Handle("/admin/user/{uid:[0-9]+}", chainAdmin(viewAdminUpdateUser))
	router.Handle("/admin/user/delete/{uid:[0-9]+}", chainAdmin(viewAdminDeleteUser))

	router.Handle("/admin/repositories", chainAdmin(viewAdminListRepositories))
	router.Handle("/admin/repository/create", chainAdmin(viewAdminCreateRepository))
	router.Handle("/admin/repository/{rid:[0-9]+}", chainAdmin(viewAdminUpdateRepository))
	router.Handle("/admin/repository/delete/{rid:[0-9]+}", chainAdmin(viewAdminDeleteRepository))

	router.Handle("/admin/mailserver/update", chainAdmin(viewAdminUpdateMailserver))

	router.Handle("/admin/config/update", chainAdmin(viewAdminUpdateConfig))

	router.Handle("/admin/repository/{rid:[0-9]+}/notification/create", chainAdmin(viewAdminCreateNotification))
	router.Handle("/admin/repository/{rid:[0-9]+}/notification/{nid:[0-9]+}", chainAdmin(viewAdminUpdateNotification))
	router.Handle("/admin/repository/{rid:[0-9]+}/notification/delete/{nid:[0-9]+}", chainAdmin(viewAdminDeleteNotification))

	router.Handle("/admin/repository/{rid:[0-9]+}/command/create", chainAdmin(viewAdminCreateCommand))
	router.Handle("/admin/repository/{rid:[0-9]+}/command/{cid:[0-9]+}", chainAdmin(viewAdminUpdateCommand))
	router.Handle("/admin/repository/{rid:[0-9]+}/command/delete/{cid:[0-9]+}", chainAdmin(viewAdminDeleteCommand))

	// add rice box to serve static files. Do not use the full middleware stack but
	// only the logging handler. We do not want "notConfigured" to run e.x. so we
	// can make the setup look nice.
	box := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	router.Handle("/static/{path:.*}", middlewareLogging(staticFileServer))

	router.Handle("/websocket", middlewareAccessKey(websocket.GetHandler()))

	return router
}
