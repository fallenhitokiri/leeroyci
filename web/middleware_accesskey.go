package web

import (
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"
)

// middlewareAccesskey tries to get a user by the access key in the
// request header. If this is not possible we return a 401 error
func middlewareAccessKey(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		header := r.Header["Accesskey"]

		if len(header) != 1 {
			http.Error(w, "Wrong access key number", http.StatusBadRequest)
			return
		}

		accessKey := header[0]

		if accessKey == "" {
			http.Error(w, "No access key", http.StatusBadRequest)
			return
		}

		_, err := database.GetUserByAccessKey(accessKey)

		if err != nil {
			http.Error(w, "Access key not found", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
