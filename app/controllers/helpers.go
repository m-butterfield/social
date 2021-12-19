package controllers

import (
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
)

func Run(port string) error {
	var err error
	if ds, err = data.Connect(); err != nil {
		return err
	}
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

type basePage struct {
	User          *data.User
	ImagesBaseURL string
	Year          string
}

func makeBasePage(c *gin.Context) *basePage {
	//user, _ := c.Get("user")
	return &basePage{
		User:          nil,
		ImagesBaseURL: lib.ImagesBaseURL,
		Year:          time.Now().Format("2006"),
	}
}

type userLoginRequest struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

func cookieLogin(w http.ResponseWriter, user *data.User) error {
	token, err := ds.CreateAccessToken(user)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "SessionToken",
		Value:   token.ID,
		Expires: token.ExpiresAt,
	})
	return nil
}
