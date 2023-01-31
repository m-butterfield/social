package data

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestCreatePost(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{Username: "testUser"}
	if err = s.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	post := &Post{
		UserID: user.ID,
	}
	if err = s.CreatePost(post); err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, post.CreatedAt)
	assert.Nil(t, post.PublishedAt)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Post{})
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
}

func TestGetPost(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{Username: "testUser"}
	if err = s.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	post := &Post{
		UserID: user.ID,
		User:   user,
	}
	if err = s.CreatePost(post); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetPost(post.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *post, *result)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Post{})
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
}

func TestPublishPost(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{Username: "testUser"}
	if err = s.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	post := &Post{
		UserID: user.ID,
	}
	if err = s.CreatePost(post); err != nil {
		t.Fatal(err)
	}
	image, err := s.GetOrCreateImage("blerp.jpg", 100, 200)
	if err != nil {
		t.Fatal(err)
	}

	if err = s.PublishPost(post.ID, []*Image{image}); err != nil {
		t.Fatal(err)
	}
	if tx := s.db.Preload("PostImages").First(&post, "id = $1", post.ID); tx.Error != nil {
		t.Fatal(tx.Error)
	}
	assert.NotNil(t, post.PublishedAt)
	assert.Equal(t, 1, len(post.PostImages))

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&PostImage{})
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Image{})
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Post{})
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
}

func TestGetUserPosts(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{Username: "testUser"}
	if err = s.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	post := &Post{
		UserID:      user.ID,
		PublishedAt: &now,
		PostImages:  []*PostImage{},
	}
	if err = s.CreatePost(post); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetUserPosts(user.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(result))
	assert.Equal(t, post.ID, result[0].ID)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Post{})
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
}

func TestUnpublishPost(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{Username: "testUser"}
	if err = s.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	post := &Post{
		PublishedAt: &now,
		UserID:      user.ID,
	}
	if err = s.CreatePost(post); err != nil {
		t.Fatal(err)
	}

	if err = s.UnpublishPost(post.ID); err != nil {
		t.Fatal(err)
	}
	if post, err = s.GetPost(post.ID); err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, post.PublishedAt)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Post{})
}
