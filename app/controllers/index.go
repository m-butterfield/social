package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func index(c *gin.Context) {
	body, err := templateRender("index", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Render(200, body)
}
