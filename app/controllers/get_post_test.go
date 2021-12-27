package controllers

import (
	"encoding/json"
	"github.com/m-butterfield/social/app/data"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPost(t *testing.T) {
	w := httptest.NewRecorder()

	expectedPost := &data.Post{
		ID: 123,
	}
	ts := &testStore{
		getPost: func(id int) (*data.Post, error) {
			return expectedPost, nil
		},
		getAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{User: expectedPost.User}, nil
		},
	}
	ds = ts

	req, err := http.NewRequest("GET", "/api/post/123", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: sessionTokenName})
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, ts.getPostCallCount)
	respBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	result := &data.Post{}
	if err = json.Unmarshal(respBody, result); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *result, *expectedPost)
}

func TestGetPostDoesntExist(t *testing.T) {
	w := httptest.NewRecorder()

	expectedPost := &data.Post{
		ID: 123,
	}
	ts := &testStore{
		getPost: func(id int) (*data.Post, error) {
			return nil, nil
		},
		getAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{User: expectedPost.User}, nil
		},
	}
	ds = ts

	req, err := http.NewRequest("GET", "/api/post/123", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: sessionTokenName})
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetPostBadID(t *testing.T) {
	w := httptest.NewRecorder()

	expectedPost := &data.Post{
		ID: 123,
	}
	ts := &testStore{
		getAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{User: expectedPost.User}, nil
		},
	}
	ds = ts

	req, err := http.NewRequest("GET", "/api/post/abc123", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: sessionTokenName})
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
