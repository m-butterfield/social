package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"log"
	"net/http"
)

type userPage struct {
	*basePage
	PageUser *data.User
	Posts    []*data.Post
}

func userIndex(c *gin.Context) {
	user, err := ds.GetUser(c.Param("userID"))
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	if user == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	posts, err := ds.GetUserPosts(user.ID)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	body, err := templateRender("app/user/index", &userPage{
		basePage: makeBasePage(c),
		PageUser: user,
		Posts:    posts,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Render(200, body)
}
