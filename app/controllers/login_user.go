package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/lib"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type userLoginRequest struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

func loginUser(c *gin.Context) {
	loginReq := &userLoginRequest{}
	err := c.Bind(loginReq)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	user, err := ds.GetUser(loginReq.UserID)
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

	token, err := ds.CreateAccessToken(user)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "SessionToken",
		Value:   token.ID,
		Expires: token.ExpiresAt,
	})
	c.Redirect(http.StatusFound, "/")
}
