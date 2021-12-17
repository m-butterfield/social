package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/static"
	"net/http"
)

func router() *gin.Engine {
	r := gin.Default()

	staticFS := http.FS(static.FS{})
	fileServer := http.FileServer(staticFS)
	addStaticHandler(r, "/css", staticFS, fileServer)
	addStaticHandler(r, "/js", staticFS, fileServer)

	r.GET("/", index)
	r.GET("/favicon.ico", favicon)

	return r
}
