package main

import (
	"net/http"

	"leeroy/database"
	"leeroy/web"
)

func main() {
	database.NewDatabase()

	router := web.Routes()
	http.ListenAndServe(":8000", router)
}
