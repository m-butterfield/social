package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"path"
)

func Run(port string) error {
	r := gin.Default()

	httpFS := http.FS(fs)
	fileServer := http.FileServer(httpFS)
	addStaticHandler(r, "/css", httpFS, fileServer)
	addStaticHandler(r, "/js", httpFS, fileServer)

	r.GET("/", index)

	return r.Run(net.JoinHostPort("", port))
}

func addStaticHandler(r *gin.Engine, prefix string, fs http.FileSystem, fileServer http.Handler) {
	handler := func(c *gin.Context) {
		f, err := fs.Open(c.Request.URL.Path)
		if err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			return
		}
		if err = f.Close(); err != nil {
			log.Printf("Error closing file %s\n", err)
		}
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
	pattern := path.Join(prefix, "/*filepath")
	r.GET(pattern, handler)
	r.HEAD(pattern, handler)
}
