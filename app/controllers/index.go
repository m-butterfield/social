package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/lib"
)

func index(c *gin.Context) {
	body, err := templateRender("index", makeBasePage(c))
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.Render(200, body)
}
