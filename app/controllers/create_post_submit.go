package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"net/http"
)

type createPostRequest struct {
	Body   string   `json:"body"`
	Images []string `json:"images"`
}

func createPostSubmit(c *gin.Context) {
	createReq := &createPostRequest{}
	err := c.Bind(createReq)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	if len(createReq.Images) == 0 {
		c.String(http.StatusBadRequest, "Please provide an image")
		return
	}
	if len(createReq.Body) > 4096 {
		c.String(http.StatusBadRequest, "Post body too long (max 4096 characters)")
		return
	}

	post := &data.Post{
		Body:   createReq.Body,
		UserID: loggedInUser(c).ID,
	}
	if err = ds.CreatePost(post); err != nil {
		lib.InternalError(err, c)
		return
	}

	if _, err := tc.CreateTask("publish_post", "social-publish-post", &lib.PublishPostRequest{
		PostID: post.ID,
		Images: createReq.Images,
	}); err != nil {
		lib.InternalError(err, c)
		return
	}

	c.JSON(http.StatusCreated, post)
}
