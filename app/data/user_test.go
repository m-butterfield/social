package data

import "testing"

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
