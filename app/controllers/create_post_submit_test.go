package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreatePostSubmit(t *testing.T) {
	w := httptest.NewRecorder()

	expectedPost := &data.Post{
		ID:   123,
		Body: "post body",
	}
	expectedImages := []string{"test.jpg"}
	ts := &testStore{
		createPost: func(post *data.Post) error {
			post.ID = 123
			assert.Equal(t, *expectedPost, *post)
			return nil
		},
		getAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{User: expectedPost.User}, nil
		},
	}
	ds = ts
	ttc := &testTaskCreator{createTask: func(taskName string, queueName string, body interface{}) (*tasks.Task, error) {
		assert.Equal(t, taskName, "publish_post")
		assert.Equal(t, queueName, "social-publish-post")
		assert.Equal(t, *body.(*lib.PublishPostRequest), lib.PublishPostRequest{
			PostID: expectedPost.ID,
			Images: expectedImages,
		})
		return nil, nil
	}}
	tc = ttc

	body, err := json.Marshal(&createPostRequest{
		Body:   expectedPost.Body,
		Images: expectedImages,
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/create_post", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: sessionTokenName})
	req.Header.Add("Content-Type", "application/json")
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, 1, ts.createPostCallCount)
	respBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	result := &data.Post{}
	if err = json.Unmarshal(respBody, result); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result.ID, expectedPost.ID)
	assert.Equal(t, result.PublishedAt, expectedPost.PublishedAt)
	assert.Equal(t, ttc.createTaskCallCount, 1)
}

func TestCreatePostSubmitNoBody(t *testing.T) {
	w := httptest.NewRecorder()

	ts := &testStore{
		createPost: func(post *data.Post) error { return nil },
		getAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{}, nil
		},
	}
	ds = ts
	ttc := &testTaskCreator{createTask: func(s string, s2 string, i interface{}) (*tasks.Task, error) {
		return nil, nil
	}}
	tc = ttc

	body, err := json.Marshal(&createPostRequest{
		Images: []string{"test.jpg"},
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/create_post", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: sessionTokenName})
	req.Header.Add("Content-Type", "application/json")
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
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
		Body:   strings.Repeat("a", 4097),
		Images: []string{"test.jpg"},
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/create_post", bytes.NewReader(body))
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
