package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"log"
)

func testRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r, err := router()
	if err != nil {
		log.Fatal(err)
	}
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
	getPost                    func(int) (*data.Post, error)
	getPostCallCount           int
	getUserPosts               func(string) ([]*data.Post, error)
	getUserPostsCallCount      int
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

func (t *testStore) GetPost(id int) (*data.Post, error) {
	t.getPostCallCount += 1
	return t.getPost(id)
}

func (t *testStore) GetOrCreateImage(string, int, int) (*data.Image, error) {
	panic("should not be called")
}

func (t *testStore) PublishPost(int, []*data.Image) error {
	panic("should not be called")
}

func (t *testStore) GetUserPosts(id string) ([]*data.Post, error) {
	t.getUserPostsCallCount += 1
	return t.getUserPosts(id)
}

type testTaskCreator struct {
	createTask          func(string, string, interface{}) (*tasks.Task, error)
	createTaskCallCount int
}

func (t *testTaskCreator) CreateTask(taskName, queueID string, body interface{}) (*tasks.Task, error) {
	t.createTaskCallCount += 1
	return t.createTask(taskName, queueID, body)
}
