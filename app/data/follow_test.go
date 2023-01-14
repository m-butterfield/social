package data

import (
	"gorm.io/gorm"
	"testing"
)

func TestCreateFollow(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{Username: "user"}
	if err = s.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	follower := &User{Username: "follower"}
	if err = s.CreateUser(follower); err != nil {
		t.Fatal(err)
	}
	follow := &Follow{
		UserID:     user.ID,
		FollowerID: follower.ID,
	}
	if err = s.CreateFollow(follow); err != nil {
		t.Fatal(err)
	}
	if user, err = s.GetUser(user.Username); err != nil {
		t.Fatal(err)
	}
	if len(user.Followers) != 1 {
		t.Error("Unexpected followers count")
	}
	if follower, err = s.GetUser(follower.Username); err != nil {
		t.Fatal(err)
	}
	if len(follower.Following) != 1 {
		t.Error("Unexpected followers count")
	}

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Follow{})
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
}

func TestGetUserFollows(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	user := &User{Username: "user"}
	if err = s.CreateUser(user); err != nil {
		t.Fatal(err)
	}
	follower := &User{Username: "follower"}
	if err = s.CreateUser(follower); err != nil {
		t.Fatal(err)
	}
	follow := &Follow{
		UserID:     user.ID,
		FollowerID: follower.ID,
	}
	if err = s.CreateFollow(follow); err != nil {
		t.Fatal(err)
	}

	follows, err := s.GetUserFollows(user.ID)
	if err != nil {
		t.Fatal(err)
	}

	if len(follows) != 1 {
		t.Error("Unexpected follows count")
	}

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Follow{})
	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
}
