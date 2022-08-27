package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserInContext(t *testing.T) {
	w := httptest.NewRecorder()
	user := &data.User{}
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil)
	c.Set(lib.UserContextKey, user)

	authRequired(c)

	assert.Equal(t, w.Result().StatusCode, 200)
}

func TestNoUserInContext(t *testing.T) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/", nil)
	authRequired(c)

	assert.Equal(t, 302, w.Result().StatusCode)
}
