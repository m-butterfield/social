package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/m-butterfield/social/app/data"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreatePostSubmit(t *testing.T) {
	w := httptest.NewRecorder()

	expectedPost := &data.Post{
		Body: "post body",
	}
	ts := &testStore{
		createPost: func(post *data.Post) error {
			assert.Equal(t, *expectedPost, *post)
			return nil
		},
		getAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{User: expectedPost.User}, nil
		},
	}
	ds = ts

	body, err := json.Marshal(&createPostRequest{
		Body: expectedPost.Body,
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/user/create_post", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: sessionTokenName})
	req.Header.Add("Content-Type", "application/json")
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, 1, ts.createPostCallCount)
}

func TestCreatePostSubmitNoBody(t *testing.T) {
	w := httptest.NewRecorder()

	ts := &testStore{
		getAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{}, nil
		},
	}
	ds = ts

	body, err := json.Marshal(&createPostRequest{})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/user/create_post", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: sessionTokenName})
	req.Header.Add("Content-Type", "application/json")
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	respBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Please provide a post body", string(respBody))
}

func TestCreatePostSubmitBodyTooLong(t *testing.T) {
	w := httptest.NewRecorder()

	ts := &testStore{
		getAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{}, nil
		},
	}
	ds = ts

	body, err := json.Marshal(&createPostRequest{
		Body: strings.Repeat("a", 4097),
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/user/create_post", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: sessionTokenName})
	req.Header.Add("Content-Type", "application/json")
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	respBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Post body too long (max 4096 characters)", string(respBody))
}
