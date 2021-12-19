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
	"time"
)

func TestCreateUser(t *testing.T) {
	w := httptest.NewRecorder()
	expectedUser := &data.User{
		ID: "test-user",
	}
	tokenID := "12345"
	expiresAt := time.Now().UTC().Add(10 * time.Minute)
	ts := &testStore{
		getUser: func(string) (*data.User, error) { return nil, nil },
		createUser: func(user *data.User) error {
			assert.Equal(t, len(user.Password), 60)
			user.Password = ""
			if *user != *expectedUser {
				t.Error("Unexpected user")
			}
			return nil
		},
		createAccessToken: func(user *data.User) (*data.AccessToken, error) {
			if *user != *expectedUser {
				t.Error("Unexpected user")
			}
			return &data.AccessToken{ID: tokenID, ExpiresAt: expiresAt}, nil
		},
	}
	ds = ts

	body, err := json.Marshal(&userLoginRequest{
		UserID:   expectedUser.ID,
		Password: "password",
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/create_user", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, ts.createUserCallCount, 1)
	assert.Equal(t, ts.createAccessTokenCallCount, 1)
	cookies := w.Result().Cookies()
	assert.Equal(t, len(cookies), 1)
	sessionCookie := cookies[0]
	assert.Equal(t, sessionCookie.Value, tokenID)
	assert.Equal(t, sessionCookie.Expires, expiresAt.Truncate(time.Second))
}

func TestCreateUserBlank(t *testing.T) {
	w := httptest.NewRecorder()
	body, err := json.Marshal(&userLoginRequest{
		UserID: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/create_user", bytes.NewReader(body))
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
	assert.Equal(t, "Please provide a user ID", string(respBody))
}

func TestCreateUserTooLong(t *testing.T) {
	w := httptest.NewRecorder()
	body, err := json.Marshal(&userLoginRequest{
		UserID: strings.Repeat("a", 65),
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/create_user", bytes.NewReader(body))
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
	assert.Equal(t, "User ID must be less than 64 characters long", string(respBody))
}

func TestCreateUserPasswordTooShort(t *testing.T) {
	w := httptest.NewRecorder()
	expectedUser := &data.User{
		ID:       "test-user",
		Password: "pass",
	}
	ts := &testStore{getUser: func(string) (*data.User, error) { return nil, nil }}
	ds = ts

	body, err := json.Marshal(&userLoginRequest{
		UserID:   expectedUser.ID,
		Password: expectedUser.Password,
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/create_user", bytes.NewReader(body))
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
	assert.Equal(t, "Password must be at least 8 characters long", string(respBody))
}

func TestCreateUserPasswordTooLong(t *testing.T) {
	w := httptest.NewRecorder()
	body, err := json.Marshal(&userLoginRequest{
		UserID:   "test-user",
		Password: strings.Repeat("a", 65),
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/create_user", bytes.NewReader(body))
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
	assert.Equal(t, "Password must be less than 64 characters long", string(respBody))
}

func TestCreateUserIDTaken(t *testing.T) {
	w := httptest.NewRecorder()
	expectedUser := &data.User{
		ID:       "test-user",
		Password: "password",
	}
	ts := &testStore{
		getUser: func(id string) (*data.User, error) {
			if id != expectedUser.ID {
				t.Error("Unexpected ID")
			}
			return &data.User{}, nil
		},
	}
	ds = ts

	body, err := json.Marshal(&userLoginRequest{
		UserID:   expectedUser.ID,
		Password: expectedUser.Password,
	})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/create_user", bytes.NewReader(body))
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
	assert.Equal(t, "Username already exists", string(respBody))
	assert.Equal(t, ts.getUserCallCount, 1)
}
