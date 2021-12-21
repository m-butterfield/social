package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"log"
)

type homePage struct {
	*basePage
	Posts []*data.Post
}

func index(c *gin.Context) {
	posts, err := ds.GetPosts()
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	body, err := templateRender("index", &homePage{
		basePage: makeBasePage(c),
		Posts:    posts,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Render(200, body)
}
