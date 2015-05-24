package web

import (
	"net/http"
	"strings"

	"leeroy/database"

	"github.com/gin-gonic/gin"
)

func notConfigured() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.String()

		if database.Configured {
			c.Next()
		} else if path == "/setup" || strings.HasPrefix(path, "/static") {
			c.Next()
		} else {
			c.Redirect(http.StatusFound, "/setup")
		}
	}
}
