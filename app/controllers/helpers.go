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
	templatePath     = "templates/"
	sessionTokenName = "SessionToken"
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

func makeBasePage(c *gin.Context) *basePage {
	return &basePage{
		User:          loggedInUser(c),
		ImagesBaseURL: lib.ImagesBaseURL,
		Year:          time.Now().UTC().Format("2006"),
	}
}

func loggedInUser(c *gin.Context) *data.User {
	result, exists := c.Get("user")
	if !exists {
		return nil
	}
	if user, ok := result.(*data.User); ok {
		return user
	}
	return nil
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
		Name:    sessionTokenName,
		Value:   token.ID,
		Expires: token.ExpiresAt,
	})
	return nil
}

func unsetSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    sessionTokenName,
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}

func getSessionCookie(c *gin.Context) (*http.Cookie, error) {
	cookie, err := c.Request.Cookie(sessionTokenName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return nil, nil
		}
		return nil, err
	}
	return cookie, nil
}
