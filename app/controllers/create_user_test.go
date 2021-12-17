package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/m-butterfield/social/app/data"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	w := httptest.NewRecorder()
	createUserCalled := false
	expectedUser := &data.User{
		ID:       "test-user",
		Password: "password",
	}
	ts := &testStore{
		getUser: func(id string) (*data.User, error) { return nil, nil },
		createUser: func(user *data.User) error {
			createUserCalled = true
			if *user != *expectedUser {
				t.Error("Unexpected user")
			}
			return nil
		},
	}
	ds = ts

	body, err := json.Marshal(map[string]interface{}{
		"userID":   expectedUser.ID,
		"password": expectedUser.Password,
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
	assert.True(t, createUserCalled)
}

func TestCreateUserIDTaken(t *testing.T) {
	w := httptest.NewRecorder()
	getUserCalled := false
	expectedUser := &data.User{
		ID:       "test-user",
		Password: "password",
	}
	ts := &testStore{
		getUser: func(id string) (*data.User, error) {
			getUserCalled = true
			if id != expectedUser.ID {
				t.Error("Unexpected ID")
			}
			return &data.User{}, nil
		},
	}
	ds = ts

	body, err := json.Marshal(map[string]interface{}{
		"userID":   expectedUser.ID,
		"password": expectedUser.Password,
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
	assert.True(t, getUserCalled)
}
