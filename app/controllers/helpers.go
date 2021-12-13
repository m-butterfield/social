package controllers

import (
	"github.com/gin-gonic/gin/render"
	"github.com/m-butterfield/social/app/static"
	"html/template"
)

const (
	templatePath = "templates/"
)

var (
	baseTemplatePaths = []string{
		templatePath + "base.gohtml",
	}
	fs static.FS
)

func templateRenderer(templateName string, data interface{}) (render.Render, error) {
	path := append([]string{templatePath + templateName + ".gohtml"}, baseTemplatePaths...)
	tmpl, err := template.ParseFS(fs, path...)
	if err != nil {
		return nil, err
	}
	return render.HTML{
		Template: tmpl,
		Data:     data,
	}, nil
}
