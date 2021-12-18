package data

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestCreateAccessToken(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{
		ID:       "testUser",
		Password: "password",
	}
	if err = s.CreateUser(user); err != nil {
		t.Fatal(err)
	}

	token, err := s.CreateAccessToken(user)
	if err != nil {
		t.Error("Unexpected error: ", err)
	}

	assert.Equal(t, token.UserID, user.ID)

	if tx := s.db.Delete(token); tx.Error != nil {
		t.Fatal(tx.Error)
	}
	if tx := s.db.Delete(user); tx.Error != nil {
		t.Fatal(tx.Error)
	}
}

func TestGetAccessToken(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{
		ID:       "testUser",
		Password: "password",
	}
	if err = s.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	expectedToken, err := s.CreateAccessToken(user)
	if err != nil {
		t.Error("Unexpected error: ", err)
	}

	token, err := s.GetAccessToken(expectedToken.ID)
	if err != nil {
		t.Error("Unexpected error")
	}

	assert.Equal(t, expectedToken.ID, token.ID)
	assert.Equal(t, *expectedToken.User, *user)

	if tx := s.db.Delete(expectedToken); tx.Error != nil {
		t.Fatal(tx.Error)
	}
	if tx := s.db.Delete(user); tx.Error != nil {
		t.Fatal(tx.Error)
	}
}

func TestGetAccessTokenInvalid(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	result, err := s.GetAccessToken("invalid")
	if result != nil {
		t.Error("Unexpected result")
	}
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
}

func TestGetAccessTokenExpired(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{
		ID:       "testUser",
		Password: "password",
	}
	err = s.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}
	token := &AccessToken{
		ID:        "1234",
		ExpiresAt: time.Now().Add(-TokenTTL),
		User:      user,
	}
	if tx := s.db.Create(token); tx.Error != nil {
		t.Fatal("Unexpected error")
	}

	result, err := s.GetAccessToken(token.ID)
	if err != nil {
		t.Fatal("Unexpected error")
	}

	assert.Nil(t, result)
	tx := s.db.First(&token, "id = $1", token.ID)
	if tx.Error == nil {
		t.Fatal("Expected NoRows Err")
	}
	if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		t.Error("Expected No Rows Err")
	}
	if tx := s.db.Delete(user); tx.Error != nil {
		t.Fatal(tx.Error)
	}
}
