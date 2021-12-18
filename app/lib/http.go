package lib

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InternalError(err error, c *gin.Context) {
	log.Println(err)
	c.AbortWithStatus(http.StatusInternalServerError)
}
