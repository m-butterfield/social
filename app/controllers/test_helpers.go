package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
)

func testRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := router()
	return r
}

type testStore struct {
	createUser                 func(*data.User) error
	createUserCallCount        int
	getUser                    func(string) (*data.User, error)
	getUserCallCount           int
	createAccessToken          func(*data.User) (*data.AccessToken, error)
	createAccessTokenCallCount int
	deleteAccessToken          func(string) error
	deleteAccessTokenCallCount int
	getAccessToken             func(string) (*data.AccessToken, error)
	getAccessTokenCallCount    int
	createPost                 func(*data.Post) error
	createPostCallCount        int
	getPosts                   func() ([]*data.Post, error)
	getPostsCallCount          int
}

func (t *testStore) CreateUser(user *data.User) error {
	t.createUserCallCount += 1
	return t.createUser(user)
}

func (t *testStore) GetUser(id string) (*data.User, error) {
	t.getUserCallCount += 1
	return t.getUser(id)
}

func (t *testStore) CreateAccessToken(user *data.User) (*data.AccessToken, error) {
	t.createAccessTokenCallCount += 1
	return t.createAccessToken(user)
}

func (t *testStore) DeleteAccessToken(id string) error {
	t.deleteAccessTokenCallCount += 1
	return t.deleteAccessToken(id)
}

func (t *testStore) GetAccessToken(id string) (*data.AccessToken, error) {
	t.getAccessTokenCallCount += 1
	return t.getAccessToken(id)
}

func (t *testStore) CreatePost(post *data.Post) error {
	t.createPostCallCount += 1
	return t.createPost(post)
}

func (t *testStore) GetPosts() ([]*data.Post, error) {
	t.getPostsCallCount += 1
	return t.getPosts()
}
