package controllers

import (
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserIndex(t *testing.T) {
	w := httptest.NewRecorder()
	testUser := &data.User{
		Username: "testUser",
	}
	req, err := http.NewRequest("GET", "/app/user/"+testUser.Username, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  lib.SessionTokenName,
		Value: "1234",
	})
	ts := &data.TestStore{
		TestGetUser: func(username string) (*data.User, error) {
			assert.Equal(t, testUser.Username, username)
			return testUser, nil
		},
		TestGetAccessToken: func(id string) (*data.AccessToken, error) {
			return &data.AccessToken{ID: "1234"}, nil
		},
		TestGetUserPosts: func(id string) ([]*data.Post, error) {
			assert.Equal(t, testUser.ID, id)
			return []*data.Post{{
				Body: "hello.",
			}}, nil
		},
	}
	ds = ts
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 1, ts.GetUserCallCount)
	assert.Equal(t, 1, ts.GetUserPostsCallCount)
}

func TestUserIndexUserDoesNotExist(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/app/user/"+"something", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  lib.SessionTokenName,
		Value: "1234",
	})
	ds = &data.TestStore{
		TestGetUser: func(username string) (*data.User, error) {
			return nil, nil
		},
		TestGetAccessToken: func(id string) (*data.AccessToken, error) {
			return &data.AccessToken{ID: "1234"}, nil
		},
	}
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
