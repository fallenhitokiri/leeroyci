package web

import (
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"

	"log"
)

type context map[string]interface{}

func NewContext(r *http.Request) context {
	var ctx = make(context)

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
