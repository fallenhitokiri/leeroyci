package main

import (
	"net/http"

	"github.com/fallenhitokiri/leeroyci/database"
	"github.com/fallenhitokiri/leeroyci/web"
)

func main() {
	database.NewDatabase()

	router := web.Routes()
	http.ListenAndServe(":8000", router)
}
