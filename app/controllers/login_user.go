package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/lib"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func loginUser(c *gin.Context) {
	loginReq := &userLoginRequest{}
	err := c.Bind(loginReq)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	user, err := ds.GetUser(loginReq.Username)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	if user == nil {
		c.String(http.StatusBadRequest, "Invalid user id")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		c.String(http.StatusBadRequest, "Invalid password")
		return
	}

	if err = cookieLogin(c.Writer, user); err != nil {
		lib.InternalError(err, c)
		return
	}
}
