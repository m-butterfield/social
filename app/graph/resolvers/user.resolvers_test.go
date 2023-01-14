package resolvers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/graph/model"
	"github.com/m-butterfield/social/app/lib"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	expectedUser := &data.User{
		Username: "test-user",
	}
	tokenID := "12345"
	expiresAt := time.Now().UTC().Add(10 * time.Minute)
	ts := &data.TestStore{
		TestGetUser: func(string) (*data.User, error) { return nil, nil },
		TestCreateUser: func(user *data.User) error {
			assert.Equal(t, len(user.Password), 60)
			user.Password = ""
			if user.ID != expectedUser.ID {
				t.Error("Unexpected user")
			}
			return nil
		},
		TestCreateAccessToken: func(user *data.User) (*data.AccessToken, error) {
			if user.ID != expectedUser.ID {
				t.Error("Unexpected user")
			}
			return &data.AccessToken{ID: tokenID, ExpiresAt: expiresAt}, nil
		},
	}
	r := Resolver{DS: ts}
	ctx := context.Background()
	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	gctx, _ := gin.CreateTestContext(w)
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)
	result, err := r.Mutation().CreateUser(ctx, model.UserCreds{
		Username: expectedUser.Username,
		Password: "password",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedUser.Username, result.Username)
	assert.Equal(t, expectedUser.Password, result.Password)
	assert.Equal(t, ts.CreateUserCallCount, 1)
	assert.Equal(t, ts.CreateAccessTokenCallCount, 1)

	cookies := w.Result().Cookies()
	assert.Equal(t, len(cookies), 1)
	sessionCookie := cookies[0]
	assert.Equal(t, sessionCookie.Value, tokenID)
	assert.Equal(t, sessionCookie.Expires, expiresAt.Truncate(time.Second))
}

func TestCreateUserBlank(t *testing.T) {
	r := Resolver{}
	_, err := r.Mutation().CreateUser(context.Background(), model.UserCreds{
		Username: " ",
		Password: "password",
	})
	if err == nil {
		t.Fatal("Expected error")
	}
	assert.Equal(t, "please provide a username", err.Error())
}

func TestCreateUserTooLong(t *testing.T) {
	r := Resolver{}
	_, err := r.Mutation().CreateUser(context.Background(), model.UserCreds{
		Username: strings.Repeat("a", 65),
		Password: "password",
	})
	if err == nil {
		t.Fatal("Expected error")
	}
	assert.Equal(t, "username must be less than 64 characters long", err.Error())
}

func TestCreateUserPasswordTooShort(t *testing.T) {
	r := Resolver{}
	_, err := r.Mutation().CreateUser(context.Background(), model.UserCreds{
		Username: "test-user",
		Password: "pass",
	})
	if err == nil {
		t.Fatal("Expected error")
	}
	assert.Equal(t, "password must be at least 8 characters long", err.Error())
}

func TestCreateUserPasswordTooLong(t *testing.T) {
	r := Resolver{}
	_, err := r.Mutation().CreateUser(context.Background(), model.UserCreds{
		Username: "test-user",
		Password: strings.Repeat("a", 65),
	})
	if err == nil {
		t.Fatal("Expected error")
	}
	assert.Equal(t, "password must be less than 64 characters long", err.Error())
}

func TestCreateUserIDTaken(t *testing.T) {
	expectedUser := &data.User{
		Username: "test-user",
		Password: "password",
	}
	ts := &data.TestStore{
		TestGetUser: func(username string) (*data.User, error) {
			if username != expectedUser.Username {
				t.Error("Unexpected ID")
			}
			return &data.User{}, nil
		},
	}
	r := Resolver{DS: ts}
	ctx := context.Background()
	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)
	gctx, _ := gin.CreateTestContext(w)
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)
	result, err := r.Mutation().CreateUser(ctx, model.UserCreds{
		Username: expectedUser.Username,
		Password: "password",
	})
	assert.Nil(t, result)
	assert.Equal(t, "username already exists", err.Error())
}

func TestLoginUser(t *testing.T) {
	w := httptest.NewRecorder()
	hashedPW, err := bcrypt.GenerateFromPassword([]byte("password"), 8)
	if err != nil {
		t.Fatal(err)
	}
	expectedUser := &data.User{
		Username: "test-user",
		Password: string(hashedPW),
	}
	tokenID := "12345"
	expiresAt := time.Now().UTC().Add(10 * time.Minute)
	ts := &data.TestStore{
		TestGetUser: func(username string) (*data.User, error) {
			assert.Equal(t, expectedUser.Username, username)
			return expectedUser, nil
		},
		TestCreateUser: func(*data.User) error { return nil },
		TestCreateAccessToken: func(user *data.User) (*data.AccessToken, error) {
			if user.ID != expectedUser.ID {
				t.Error("Unexpected user")
			}
			return &data.AccessToken{ID: tokenID, ExpiresAt: expiresAt}, nil
		},
	}
	r := Resolver{DS: ts}

	ctx := context.Background()
	gin.SetMode(gin.ReleaseMode)
	gctx, _ := gin.CreateTestContext(w)
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)
	result, err := r.Mutation().Login(ctx, model.UserCreds{
		Username: " " + expectedUser.Username,
		Password: "password",
	})
	assert.NotNil(t, result)
	cookies := w.Result().Cookies()
	sessionCookie := cookies[0]
	assert.Equal(t, len(cookies), 1)
	assert.Equal(t, sessionCookie.Value, tokenID)
	assert.Equal(t, sessionCookie.Expires, expiresAt.Truncate(time.Second))
	assert.Equal(t, ts.GetUserCallCount, 1)
	assert.Equal(t, ts.CreateAccessTokenCallCount, 1)
}

func TestLoginBadUserID(t *testing.T) {
	w := httptest.NewRecorder()
	ts := &data.TestStore{
		TestGetUser: func(username string) (*data.User, error) {
			return nil, nil
		},
		TestCreateUser: func(*data.User) error { return nil },
	}
	r := Resolver{DS: ts}
	ctx := context.Background()
	gctx, _ := gin.CreateTestContext(w)
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)
	result, err := r.Mutation().Login(ctx, model.UserCreds{
		Username: "blah",
		Password: "password",
	})
	assert.Nil(t, result)
	assert.Equal(t, "invalid username", err.Error())
}

func TestLoginBadPassword(t *testing.T) {
	w := httptest.NewRecorder()
	hashedPW, err := bcrypt.GenerateFromPassword([]byte("password"), 8)
	if err != nil {
		t.Fatal(err)
	}
	expectedUser := &data.User{
		Username: "test-user",
		Password: string(hashedPW),
	}
	ts := &data.TestStore{
		TestGetUser: func(username string) (*data.User, error) {
			assert.Equal(t, expectedUser.Username, username)
			return expectedUser, nil
		},
		TestCreateUser: func(*data.User) error { return nil },
	}
	r := Resolver{DS: ts}

	ctx := context.Background()
	gctx, _ := gin.CreateTestContext(w)
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)
	result, err := r.Mutation().Login(ctx, model.UserCreds{
		Username: "test-user",
		Password: "badpassword",
	})
	assert.Nil(t, result)
	assert.Equal(t, "invalid password", err.Error())
}

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	tokenID := "1234"
	ts := &data.TestStore{
		TestGetAccessToken: func(s string) (*data.AccessToken, error) {
			return &data.AccessToken{ID: tokenID}, nil
		},
		TestDeleteAccessToken: func(id string) error {
			assert.Equal(t, id, tokenID)
			return nil
		},
	}
	r := Resolver{DS: ts}

	ctx := context.Background()
	gctx, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("POST", "/graphql/query", nil)
	gctx.Request = req
	gctx.Request.AddCookie(&http.Cookie{
		Name:  lib.SessionTokenName,
		Value: tokenID,
	})
	ctx = context.WithValue(ctx, lib.GinContextKey, gctx)
	result, err := r.Mutation().Logout(ctx)
	assert.Nil(t, err)
	assert.True(t, result)

	assert.Equal(t, ts.DeleteAccessTokenCallCount, 1)
	cookies := w.Result().Cookies()
	assert.Equal(t, len(cookies), 1)
	sessionCookie := cookies[0]
	assert.Equal(t, sessionCookie.Value, "")
	assert.Equal(t, sessionCookie.Expires, time.Unix(0, 0).UTC())
}
