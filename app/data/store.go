package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

func Connect() (Store, error) {
	s, err := getDS()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func getDS() (*ds, error) {
	var logLevel logger.LogLevel
	if os.Getenv("SQL_LOGS") == "true" {
		logLevel = logger.Info
	} else {
		logLevel = logger.Silent
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_SOCKET")), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}
	return &ds{db: db}, nil
}

type Store interface {
	CreateUser(*User) error
	CreateFollow(*Follow) error
	DeleteFollow(*Follow) error
	GetUser(string) (*User, error)
	GetUserFollows(string) ([]*Follow, error)
	CreateAccessToken(*User) (*AccessToken, error)
	DeleteAccessToken(string) error
	GetAccessToken(string) (*AccessToken, error)
	CreatePost(*Post) error
	GetPosts(*time.Time) ([]*Post, error)
	GetPost(string) (*Post, error)
	GetOrCreateImage(string, int, int) (*Image, error)
	PublishPost(string, []*Image) error
	UnpublishPost(string) error
	GetUserPosts(string, *time.Time) ([]*Post, error)
	GetUsersPosts([]string, *time.Time) ([]*Post, error)
	CreateComment(*Comment) error
	GetComments(string, *time.Time) ([]*Comment, error)
}

type ds struct {
	db *gorm.DB
}
