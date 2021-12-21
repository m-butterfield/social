package data

import (
	"testing"
)

func TestCreatePost(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	post := &Post{
		Body: "test body",
		User: &User{ID: "testUser"},
	}
	if err = s.CreatePost(post); err != nil {
		t.Fatal(err)
	}

	if tx := s.db.Delete(post); tx.Error != nil {
		t.Fatal(tx.Error)
	}
}
