package controllers

import (
	"github.com/m-butterfield/social/app/data"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserIndex(t *testing.T) {
	w := httptest.NewRecorder()
	testUser := &data.User{
		ID: "testUser",
	}
	req, err := http.NewRequest("GET", "/app/user/"+testUser.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  sessionTokenName,
		Value: "1234",
	})
	ts := &testStore{
		getUser: func(id string) (*data.User, error) {
			assert.Equal(t, testUser.ID, id)
			return testUser, nil
		},
		getAccessToken: func(id string) (*data.AccessToken, error) {
			return &data.AccessToken{ID: "1234"}, nil
		},
		getUserPosts: func(id string) ([]*data.Post, error) {
			assert.Equal(t, testUser.ID, id)
			return []*data.Post{{
				Body: "hello.",
			}}, nil
		},
	}
	ds = ts
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, ts.getUserCallCount)
	assert.Equal(t, 1, ts.getUserPostsCallCount)
}

func TestUserIndexUserDoesNotExist(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/app/user/"+"something", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  sessionTokenName,
		Value: "1234",
	})
	ds = &testStore{
		getUser: func(id string) (*data.User, error) {
			return nil, nil
		},
		getAccessToken: func(id string) (*data.AccessToken, error) {
			return &data.AccessToken{ID: "1234"}, nil
		},
	}
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
