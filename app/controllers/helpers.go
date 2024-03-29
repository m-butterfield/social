package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"github.com/m-butterfield/social/app/static"
	"html/template"
	"net"
	"net/http"
	"time"
)

const (
	templatePath = "templates/"
)

var (
	baseTemplatePaths = []string{
		templatePath + "base.gohtml",
	}
	ds data.Store
	tc lib.TaskCreator
)

func Run(port string) error {
	var err error
	if ds, err = data.Connect(); err != nil {
		return err
	}
	if tc, err = lib.NewTaskCreator(); err != nil {
		return err
	}
	r, err := router()
	if err != nil {
		return err
	}
	return r.Run(net.JoinHostPort("", port))
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

type basePage struct {
	User          *data.User
	ImagesBaseURL string
	Year          string
}

func makeBasePage() *basePage {
	return &basePage{
		ImagesBaseURL: lib.ImagesBaseURL,
		Year:          time.Now().UTC().Format("2006"),
	}
}

func unsetSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    lib.SessionTokenName,
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}

func getSessionCookie(c *gin.Context) (*http.Cookie, error) {
	cookie, err := c.Request.Cookie(lib.SessionTokenName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return nil, nil
		}
		return nil, err
	}
	return cookie, nil
}
