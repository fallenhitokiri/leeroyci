package web

import (
	"log"
	"net/http"
	"strings"

	"leeroy/database"

	"github.com/gin-gonic/gin"
)

func notConfigured() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := database.GetConfig()
		path := c.Request.URL.String()

		log.Println(path)

		if cfg.ID != 0 {
			log.Println("cfg exists")
			c.Next()
		} else if path == "/setup" || strings.HasPrefix(path, "/static") {
			log.Println("setup or static")
			c.Next()
		} else {
			log.Println("redirect")
			c.Redirect(http.StatusFound, "/setup")
		}
	}
}
