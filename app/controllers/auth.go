package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/lib"
)

func auth(c *gin.Context) {
	cookie, err := getSessionCookie(c)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	if cookie == nil {
		return
	}
	token, err := ds.GetAccessToken(cookie.Value)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	if token == nil {
		unsetSessionCookie(c.Writer)
		return
	}
	c.Set(lib.UserContextKey, token.User)
}
