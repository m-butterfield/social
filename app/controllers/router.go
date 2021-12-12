package controllers

import (
	"github.com/gin-gonic/gin"
	"net"
)

func Run(port string) error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r.Run(net.JoinHostPort("", port))
}
