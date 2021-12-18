package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/lib"
	"net/http"
)

func auth(c *gin.Context) {
	id, err := c.Cookie("SessionToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return
		}
		lib.InternalError(err, c)
		return
	}
	token, err := ds.GetAccessToken(id)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	if token != nil {
		c.Set("user", token.User)
		return
	}
}
