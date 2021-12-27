package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/lib"
)

func router() (*gin.Engine, error) {
	r, err := lib.BaseRouter()
	if err != nil {
		return nil, err
	}

	r.POST("/publish_post", publishPost)

	return r, nil
}
