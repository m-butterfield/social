package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/m-butterfield/social/app/data"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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
	ts := &testStore{
		getUser: func(username string) (*data.User, error) {
			assert.Equal(t, expectedUser.Username, username)
			return expectedUser, nil
		},
		createUser: func(*data.User) error { return nil },
		createAccessToken: func(user *data.User) (*data.AccessToken, error) {
			if *user != *expectedUser {
				t.Error("Unexpected user")
			}
			return &data.AccessToken{ID: tokenID, ExpiresAt: expiresAt}, nil
		},
	}
	ds = ts

	body, err := json.Marshal(&userLoginRequest{
		Username: expectedUser.Username,
		Password: "password",
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, ts.getUserCallCount, 1)
	assert.Equal(t, ts.createAccessTokenCallCount, 1)
	cookies := w.Result().Cookies()
	assert.Equal(t, len(cookies), 1)
	sessionCookie := cookies[0]
	assert.Equal(t, sessionCookie.Value, tokenID)
	assert.Equal(t, sessionCookie.Expires, expiresAt.Truncate(time.Second))
}

func TestLoginBadUserID(t *testing.T) {
	w := httptest.NewRecorder()
	ts := &testStore{
		getUser: func(username string) (*data.User, error) {
			return nil, nil
		},
		createUser: func(*data.User) error { return nil },
	}
	ds = ts

	body, err := json.Marshal(&userLoginRequest{
		Username: "blah",
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	respBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Invalid user id", string(respBody))
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
	ts := &testStore{
		getUser: func(username string) (*data.User, error) {
			assert.Equal(t, expectedUser.Username, username)
			return expectedUser, nil
		},
		createUser: func(*data.User) error { return nil },
	}
	ds = ts

	body, err := json.Marshal(&userLoginRequest{
		Username: expectedUser.Username,
		Password: "incorrect",
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	respBody, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "Invalid password", string(respBody))
}
