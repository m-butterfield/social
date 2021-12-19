package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/lib"
	"net/http"
)

func logout(c *gin.Context) {
	cookie, err := getSessionCookie(c)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	if cookie == nil {
		return
	}
	if err := ds.DeleteAccessToken(cookie.Value); err != nil {
		lib.InternalError(err, c)
		return
	}
	unsetSessionCookie(c.Writer)
	c.Redirect(http.StatusFound, "/")
}
