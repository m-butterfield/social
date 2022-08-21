package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func createUser(c *gin.Context) {
	createReq := &userLoginRequest{}
	err := c.Bind(createReq)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	if createReq.Username == "" {
		c.String(http.StatusBadRequest, "Please provide a user Username")
		return
	}
	if len(createReq.Username) > 64 {
		c.String(http.StatusBadRequest, "User Username must be less than 64 characters long")
		return
	}
	if len(createReq.Password) < 8 {
		c.String(http.StatusBadRequest, "Password must be at least 8 characters long")
		return
	}
	if len(createReq.Password) > 64 {
		c.String(http.StatusBadRequest, "Password must be less than 64 characters long")
		return
	}

	if exising, err := ds.GetUser(createReq.Username); err != nil {
		lib.InternalError(err, c)
		return
	} else if exising != nil {
		c.String(http.StatusBadRequest, "Username already exists")
		return
	}

	hashedPW, err := bcrypt.GenerateFromPassword([]byte(createReq.Password), 8)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	user := &data.User{
		Username: createReq.Username,
		Password: string(hashedPW),
	}
	if err = ds.CreateUser(user); err != nil {
		lib.InternalError(err, c)
		return
	}

	if err = cookieLogin(c.Writer, user); err != nil {
		lib.InternalError(err, c)
		return
	}
}
