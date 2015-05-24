package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupGET(c *gin.Context) {
	c.String(http.StatusOK, "cfg")
}
