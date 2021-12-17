package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/m-butterfield/social/app/lib"
	"github.com/m-butterfield/social/app/static"
	"html/template"
	"log"
	"net"
	"net/http"
	"path"
	"time"
)

const (
	templatePath = "templates/"
)

var (
	baseTemplatePaths = []string{
		templatePath + "base.gohtml",
	}
)

func Run(port string) error {
	return router().Run(net.JoinHostPort("", port))
}

func templateRender(name string, data interface{}) (render.Render, error) {
	paths := append([]string{templatePath + name + ".gohtml"}, baseTemplatePaths...)
	tmpl, err := template.ParseFS(static.FS{}, paths...)
	if err != nil {
		return nil, err
	}
	return render.HTML{
		Template: tmpl,
		Data:     data,
	}, nil
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

type basePage struct {
	ImagesBaseURL string
	Year          string
}

func makeBasePage() *basePage {
	return &basePage{
		ImagesBaseURL: lib.ImagesBaseURL,
		Year:          time.Now().Format("2006"),
	}
}
