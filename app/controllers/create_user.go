package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"net/http"
)

func createUser(c *gin.Context) {
	user := &data.User{}
	err := c.Bind(user)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	if exising, err := ds.GetUser(user.ID); err != nil {
		lib.InternalError(err, c)
		return
	} else if exising != nil {
		c.String(http.StatusBadRequest, "Username already exists")
		return
	}

	if err = ds.CreateUser(user); err != nil {
		lib.InternalError(err, c)
		return
	}

	c.Status(http.StatusCreated)
}
