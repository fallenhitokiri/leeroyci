package web

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
)

func staticServe(c *gin.Context) {
	static, err := rice.FindBox("../static")

	if err != nil {
		panic(err)
	}

	original := c.Request.URL.Path
	c.Request.URL.Path = c.Params.ByName("filepath")

	http.FileServer(static.HTTPBox()).ServeHTTP(c.Writer, c.Request)
	c.Request.URL.Path = original
}
