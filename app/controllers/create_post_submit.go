package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"net/http"
)

type createPostRequest struct {
	Body string `json:"body"`
}

func createPostSubmit(c *gin.Context) {
	createReq := &createPostRequest{}
	err := c.Bind(createReq)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	if createReq.Body == "" {
		c.String(http.StatusBadRequest, "Please provide a post body")
		return
	}
	if len(createReq.Body) > 4096 {
		c.String(http.StatusBadRequest, "Post body too long (max 4096 characters)")
		return
	}

	post := &data.Post{
		Body: createReq.Body,
		User: loggedInUser(c),
	}
	if err = ds.CreatePost(post); err != nil {
		lib.InternalError(err, c)
		return
	}

	c.Status(201)
}
