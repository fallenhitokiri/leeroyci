package main

import (
	"leeroy/database"
	"leeroy/web"

	"github.com/gin-gonic/gin"
)

func main() {
	database.NewDatabase()

	engine := gin.Default()
	web.AddRoutes(engine)
	engine.Run(":8000")
}
