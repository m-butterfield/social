package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginUser(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
