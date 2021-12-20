package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"net/http"
)

func authRequired(c *gin.Context) {
	if val, exists := c.Get("user"); exists {
		if _, ok := val.(*data.User); ok {
			return
		}
	}
	c.Redirect(http.StatusFound, "/login")
	c.Abort()
}
