package resolvers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/graph/model"
	"github.com/m-butterfield/social/app/lib"
	"github.com/stretchr/testify/assert"
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
			if *user != *expectedUser {
				t.Error("Unexpected user")
			}
			return nil
		},
		TestCreateAccessToken: func(user *data.User) (*data.AccessToken, error) {
			if *user != *expectedUser {
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
	result, err := r.Mutation().CreateUser(ctx, model.CreateUser{
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
	_, err := r.Mutation().CreateUser(context.Background(), model.CreateUser{
		Username: "",
		Password: "password",
	})
	if err == nil {
		t.Fatal("Expected error")
	}
	assert.Equal(t, "please provide a username", err.Error())
}

func TestCreateUserTooLong(t *testing.T) {
	r := Resolver{}
	_, err := r.Mutation().CreateUser(context.Background(), model.CreateUser{
		Username: strings.Repeat("a", 65),
		Password: "password",
	})
	if err == nil {
		t.Fatal("Expected error")
	}
	assert.Equal(t, "username must be less than 64 characters long", err.Error())
}

//func TestCreateUserPasswordTooShort(t *testing.T) {
//	w := httptest.NewRecorder()
//	expectedUser := &data.User{
//		Username: "test-user",
//		Password: "pass",
//	}
//	ts := &data.TestStore{
//		TestGetUser: func(string) (*data.User, error) { return nil, nil },
//	}
//	ds = ts
//
//	body, err := json.Marshal(&userLoginRequest{
//		Username: expectedUser.Username,
//		Password: expectedUser.Password,
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//	req, err := http.NewRequest("POST", "/create_user", bytes.NewReader(body))
//	if err != nil {
//		t.Fatal(err)
//	}
//	req.Header.Add("Content-Type", "application/json")
//	testRouter().ServeHTTP(w, req)
//
//	assert.Equal(t, 400, w.Code)
//	respBody, err := io.ReadAll(w.Body)
//	if err != nil {
//		t.Fatal(err)
//	}
//	assert.Equal(t, "Password must be at least 8 characters long", string(respBody))
//}
//
//func TestCreateUserPasswordTooLong(t *testing.T) {
//	w := httptest.NewRecorder()
//	body, err := json.Marshal(&userLoginRequest{
//		Username: "test-user",
//		Password: strings.Repeat("a", 65),
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//	req, err := http.NewRequest("POST", "/create_user", bytes.NewReader(body))
//	if err != nil {
//		t.Fatal(err)
//	}
//	req.Header.Add("Content-Type", "application/json")
//	testRouter().ServeHTTP(w, req)
//
//	assert.Equal(t, 400, w.Code)
//	respBody, err := io.ReadAll(w.Body)
//	if err != nil {
//		t.Fatal(err)
//	}
//	assert.Equal(t, "Password must be less than 64 characters long", string(respBody))
//}
//
//func TestCreateUserIDTaken(t *testing.T) {
//	w := httptest.NewRecorder()
//	expectedUser := &data.User{
//		Username: "test-user",
//		Password: "password",
//	}
//	ts := &data.TestStore{
//		TestGetUser: func(username string) (*data.User, error) {
//			if username != expectedUser.Username {
//				t.Error("Unexpected ID")
//			}
//			return &data.User{}, nil
//		},
//	}
//	ds = ts
//
//	body, err := json.Marshal(&userLoginRequest{
//		Username: expectedUser.Username,
//		Password: expectedUser.Password,
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//	req, err := http.NewRequest("POST", "/create_user", bytes.NewReader(body))
//	if err != nil {
//		t.Fatal(err)
//	}
//	req.Header.Add("Content-Type", "application/json")
//	testRouter().ServeHTTP(w, req)
//
//	assert.Equal(t, 400, w.Code)
//	respBody, err := io.ReadAll(w.Body)
//	if err != nil {
//		t.Fatal(err)
//	}
//	assert.Equal(t, "Username already exists", string(respBody))
//	assert.Equal(t, ts.GetUserCallCount, 1)
//}
