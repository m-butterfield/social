package data

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connect() (Store, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}
	return &ds{db: db}, nil
}

func Migrate() error {
	db, err := getDB()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(
		&User{},
	)
	if err != nil {
		return err
	}
	return nil
}

func getDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_SOCKET")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Store interface {
	CreateUser(*User) error
	GetUser(string) (*User, error)
}

type ds struct {
	db *gorm.DB
}
