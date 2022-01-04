package data

import (
	"gorm.io/gorm"
	"testing"
)

func TestGetUserInvalidUserID(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	result, err := s.GetUser("invalid")
	if result != nil {
		t.Error("Unexpected result")
	}
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
}

func TestAddFollower(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{ID: "testUser"}
	if err = s.CreateUser(user); err != nil {
		t.Fatal(err)
	}

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
}
