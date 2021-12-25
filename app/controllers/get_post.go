package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/lib"
	"net/http"
	"strconv"
)

func getPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	post, err := ds.GetPost(id)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	if post == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(200, post)
}
