package data

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Username string `gorm:"type:varchar(64);not null;index:unique"`
	Password string `gorm:"type:varchar(60);not null"`
}

func (s *ds) CreateUser(user *User) error {
	if tx := s.db.Create(user); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) GetUser(username string) (*User, error) {
	user := &User{}
	tx := s.db.First(&user, "username = $1", username)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return user, nil
}
