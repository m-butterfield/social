package data

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:uuid;default:uuid_generate_v4()" json:"-"`
	Username string `gorm:"type:citext;not null;unique" json:"username"`
	Password string `gorm:"type:varchar(60);not null" json:"password"`
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
