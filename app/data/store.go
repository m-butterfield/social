package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func Connect() (Store, error) {
	s, err := getDS()
	if err != nil {
		return nil, err
	}
	return s, nil
}

func getDS() (*ds, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_SOCKET")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return &ds{db: db}, nil
}

type Store interface {
	CreateUser(*User) error
	GetUser(string) (*User, error)
	CreateAccessToken(*User) (*AccessToken, error)
	DeleteAccessToken(string) error
	GetAccessToken(string) (*AccessToken, error)
	CreatePost(*Post) error
	GetPosts() ([]*Post, error)
}

type ds struct {
	db *gorm.DB
}
