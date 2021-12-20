package controllers

import (
	"github.com/m-butterfield/social/app/data"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePost(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/user/create_post", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  sessionTokenName,
		Value: "1234",
	})
	ds = &testStore{
		getAccessToken: func(id string) (*data.AccessToken, error) {
			return &data.AccessToken{ID: "1234"}, nil
		},
	}
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
