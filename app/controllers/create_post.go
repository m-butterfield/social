package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func createPost(c *gin.Context) {
	body, err := templateRender("app/create_post", makeBasePage())
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Render(200, body)
}
