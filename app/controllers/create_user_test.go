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
)

func TestCreateUser(t *testing.T) {
	w := httptest.NewRecorder()
	expectedUser := &data.User{
		ID: "test-user",
	}
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

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, ts.createUseCallCount, 1)
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
