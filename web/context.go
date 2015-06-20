package web

import (
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"
)

type responseContext map[string]interface{}

type requestContextKey string

const RequestContextUser = "user"

func NewContext(r *http.Request) responseContext {
	var ctx = make(responseContext)

	session, _ := store.Get(r, "leeroyci")
	session_key := session.Values["session_key"]

	if session_key != nil {
		key := session_key.(string)
		user, err := database.GetUser(key)

		if err == nil {
			ctx["user"] = user
		}
	}

	return ctx
}
