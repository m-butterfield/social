package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAuthGoodToken(t *testing.T) {
	w := httptest.NewRecorder()
	tokenID := "1234"
	user := &data.User{}
	ts := &data.TestStore{
		TestGetAccessToken: func(id string) (*data.AccessToken, error) {
			return &data.AccessToken{ID: tokenID, User: user}, nil
		},
	}
	ds = ts
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)
	var err error
	c.Request, err = http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	c.Request.AddCookie(&http.Cookie{
		Name:  sessionTokenName,
		Value: tokenID,
	})

	auth(c)

	val, exists := c.Get("user")
	assert.True(t, exists)
	assert.Equal(t, val.(*data.User), user)
}

func TestAuthBadToken(t *testing.T) {
	w := httptest.NewRecorder()
	tokenID := "1234"
	ts := &data.TestStore{
		TestGetAccessToken: func(id string) (*data.AccessToken, error) {
			return nil, nil
		},
	}
	ds = ts
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)
	var err error
	c.Request, err = http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	c.Request.AddCookie(&http.Cookie{
		Name:  sessionTokenName,
		Value: tokenID,
	})

	auth(c)

	_, exists := c.Get("user")
	assert.False(t, exists)
	cookies := w.Result().Cookies()
	assert.Equal(t, len(cookies), 1)
	sessionCookie := cookies[0]
	assert.Equal(t, sessionCookie.Value, "")
	assert.Equal(t, sessionCookie.Expires, time.Unix(0, 0).UTC())
}

func TestAuthNoToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var err error
	gin.SetMode(gin.ReleaseMode)
	c.Request, err = http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	auth(c)

	_, exists := c.Get("user")
	assert.False(t, exists)
}
