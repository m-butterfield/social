package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func login(c *gin.Context) {
	body, err := templateRender("login", makeBasePage(c))
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Render(200, body)
}
