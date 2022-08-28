package resolvers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/graph/model"
	"github.com/m-butterfield/social/app/lib"
	"github.com/m-butterfield/social/app/tasks"
	"github.com/stretchr/testify/assert"
	cloudtask "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreatePostSubmit(t *testing.T) {
	w := httptest.NewRecorder()
	expectedPost := &data.Post{
		ID:   "123",
		Body: "post body",
	}
	expectedImages := []string{"test.jpg"}
	expectedUserID := "123"
	ts := &data.TestStore{
		TestCreatePost: func(post *data.Post) error {
			post.ID = "123"
			assert.Equal(t, expectedPost.Body, post.Body)
			assert.Equal(t, expectedUserID, post.UserID)
			return nil
		},
		TestGetAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{User: &data.User{ID: expectedUserID}}, nil
		},
	}
	ttc := &tasks.TestTaskCreator{TestCreateTask: func(taskName string, queueName string, body interface{}) (*cloudtask.Task, error) {
		assert.Equal(t, taskName, "publish_post")
		assert.Equal(t, queueName, "social-publish-post")
		assert.Equal(t, *body.(*lib.PublishPostRequest), lib.PublishPostRequest{
			PostID: expectedPost.ID,
			Images: expectedImages,
		})
		return nil, nil
	}}
	r := Resolver{DS: ts, TC: ttc}

	ctx := context.Background()
	gin.SetMode(gin.ReleaseMode)
	gctx, _ := gin.CreateTestContext(w)
	gctx.Set(lib.UserContextKey, &data.User{ID: expectedUserID})
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)
	result, err := r.Mutation().CreatePost(ctx, model.CreatePostInput{
		Body:   expectedPost.Body,
		Images: expectedImages,
	})
	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, 1, ts.CreatePostCallCount)
	assert.Equal(t, result.ID, expectedPost.ID)
	assert.Equal(t, result.PublishedAt, expectedPost.PublishedAt)
	assert.Equal(t, ttc.CreateTaskCallCount, 1)
}

func TestCreatePostSubmitNoBody(t *testing.T) {
	w := httptest.NewRecorder()
	ts := &data.TestStore{
		TestCreatePost: func(post *data.Post) error { return nil },
		TestGetAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{User: &data.User{Username: "testUser"}}, nil
		},
	}
	ttc := &tasks.TestTaskCreator{TestCreateTask: func(s string, s2 string, i interface{}) (*cloudtask.Task, error) {
		return nil, nil
	}}
	r := Resolver{DS: ts, TC: ttc}

	ctx := context.Background()
	gin.SetMode(gin.ReleaseMode)
	gctx, _ := gin.CreateTestContext(w)
	gctx.Set(lib.UserContextKey, &data.User{ID: "123"})
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)
	result, err := r.Mutation().CreatePost(ctx, model.CreatePostInput{
		Images: []string{"test.jpg"},
	})
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestCreatePostSubmitBodyTooLong(t *testing.T) {
	w := httptest.NewRecorder()
	ts := &data.TestStore{
		TestCreatePost: func(post *data.Post) error { return nil },
		TestGetAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{User: &data.User{Username: "testUser"}}, nil
		},
	}
	ttc := &tasks.TestTaskCreator{TestCreateTask: func(s string, s2 string, i interface{}) (*cloudtask.Task, error) {
		return nil, nil
	}}
	r := Resolver{DS: ts, TC: ttc}

	ctx := context.Background()
	gin.SetMode(gin.ReleaseMode)
	gctx, _ := gin.CreateTestContext(w)
	gctx.Set(lib.UserContextKey, &data.User{ID: "123"})
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)
	result, err := r.Mutation().CreatePost(ctx, model.CreatePostInput{
		Body:   strings.Repeat("a", 4097),
		Images: []string{"test.jpg"},
	})
	assert.Equal(t, "post body too long (max 4096 characters)", err.Error())
	assert.Nil(t, result)
}

func TestGetPost(t *testing.T) {
	w := httptest.NewRecorder()
	expectedPost := &data.Post{
		ID: "123",
	}
	ts := &data.TestStore{
		TestGetPost: func(id string) (*data.Post, error) {
			return expectedPost, nil
		},
		TestGetAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{User: expectedPost.User}, nil
		},
	}
	r := Resolver{DS: ts}
	ctx := context.Background()
	gin.SetMode(gin.ReleaseMode)
	gctx, _ := gin.CreateTestContext(w)
	gctx.Set(lib.UserContextKey, &data.User{ID: "123"})
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)

	result, err := r.Query().GetPost(ctx, expectedPost.ID)

	assert.Nil(t, err)
	assert.Equal(t, 1, ts.GetPostCallCount)
	assert.Equal(t, *expectedPost, *result)
}

func TestGetPostDoesntExist(t *testing.T) {
	w := httptest.NewRecorder()
	expectedPost := &data.Post{
		ID: "123",
	}
	ts := &data.TestStore{
		TestGetPost: func(id string) (*data.Post, error) {
			return nil, nil
		},
		TestGetAccessToken: func(string) (*data.AccessToken, error) {
			return &data.AccessToken{User: expectedPost.User}, nil
		},
	}
	r := Resolver{DS: ts}
	ctx := context.Background()
	gin.SetMode(gin.ReleaseMode)
	gctx, _ := gin.CreateTestContext(w)
	gctx.Set(lib.UserContextKey, &data.User{ID: "123"})
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)

	result, err := r.Query().GetPost(ctx, expectedPost.ID)

	assert.Nil(t, result)
	assert.Equal(t, "post not found", err.Error())
}

func TestGetUserPosts(t *testing.T) {
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
	r := Resolver{DS: ts}

	result, err := r.Query().GetUserPosts(context.Background(), testUser.Username)

	assert.Nil(t, err)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, 1, ts.GetUserCallCount)
	assert.Equal(t, 1, ts.GetUserPostsCallCount)
}

func TestGetUserPostsDoesNotExist(t *testing.T) {
	req, err := http.NewRequest("GET", "/app/user/"+"something", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  lib.SessionTokenName,
		Value: "1234",
	})
	ts := &data.TestStore{
		TestGetUser: func(username string) (*data.User, error) {
			return nil, nil
		},
		TestGetAccessToken: func(id string) (*data.AccessToken, error) {
			return &data.AccessToken{ID: "1234"}, nil
		},
	}
	r := Resolver{DS: ts}

	result, err := r.Query().GetUserPosts(context.Background(), "none")

	assert.Nil(t, result)
	assert.Equal(t, "user not found", err.Error())
}
