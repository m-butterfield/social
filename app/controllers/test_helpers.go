package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func testRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r, err := router()
	if err != nil {
		log.Fatal(err)
	}
	return r
}
