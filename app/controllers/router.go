package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/static"
	"log"
	"net/http"
	"path"
)

func router() *gin.Engine {
	r := gin.New()
	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	r.Use(auth)

	staticFS := http.FS(static.FS{})
	fileServer := http.FileServer(staticFS)
	addStaticHandler(r, "/css", fileServer)
	addStaticHandler(r, "/js", fileServer)

	r.GET("/", index)
	r.GET("/favicon.ico", favicon)

	r.GET("/login", login)
	r.POST("/login", loginUser)
	r.POST("/create_user", createUser)

	return r
}

func addStaticHandler(r *gin.Engine, prefix string, fileServer http.Handler) {
	handler := func(c *gin.Context) { fileServer.ServeHTTP(c.Writer, c.Request) }
	pattern := path.Join(prefix, "/*filepath")
	r.GET(pattern, handler)
	r.HEAD(pattern, handler)
}
