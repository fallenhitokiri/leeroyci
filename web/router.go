package web

import (
	"github.com/gin-gonic/gin"
)

// AddRoutes adds all routes for the webinterface to a gin router.
func AddRoutes(engine *gin.Engine) {
	// setup middleware
	engine.Use(notConfigured())

	engine.GET("/setup", setupGET)
	engine.GET("/static/*filepath", staticServe)
}
