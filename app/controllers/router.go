package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/lib"
	"github.com/m-butterfield/social/app/static"
	"net/http"
	"path"
)

func router() (*gin.Engine, error) {
	r, err := lib.BaseRouter()
	if err != nil {
		return nil, err
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
	r.GET("/logout", logout)

	app := r.Group("/app")
	app.Use(authRequired)
	app.GET("/create_post", createPost)

	api := r.Group("/api")
	app.Use(authRequired)
	api.POST("/signed_upload_url", signedUploadURL)
	api.POST("/create_post", createPostSubmit)
	api.GET("/post/:ID", getPost)

	return r, nil
}

func addStaticHandler(r *gin.Engine, prefix string, fileServer http.Handler) {
	handler := func(c *gin.Context) { fileServer.ServeHTTP(c.Writer, c.Request) }
	pattern := path.Join(prefix, "/*filepath")
	r.GET(pattern, handler)
	r.HEAD(pattern, handler)
}
