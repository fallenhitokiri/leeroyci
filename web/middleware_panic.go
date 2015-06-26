package web

import (
	"errors"
	"log"
	"net/http"
)

// middlewarePanic recovers from panics during the whole request process and
// renders a 500 error page.
func middlewarePanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var err error

		defer func() {
			rec := recover()

			if rec != nil {
				switch kind := rec.(type) {
				case string:
					err = errors.New(kind)
				case error:
					err = kind
				default:
					err = errors.New("Unknown error")
				}
				log.Println("panic: ", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
