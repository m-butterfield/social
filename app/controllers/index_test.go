package controllers

import (
	"github.com/m-butterfield/social/app/data"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	w := httptest.NewRecorder()

	ts := &testStore{
		getPosts: func() ([]*data.Post, error) {
			return []*data.Post{}, nil
		},
	}
	ds = ts

	req, _ := http.NewRequest("GET", "/", nil)
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, ts.getPostsCallCount)
}
