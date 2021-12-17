package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey" json:"userID"`
	Password string `gorm:"not null" json:"password"`
}

func (d *ds) CreateUser(user *User) error {
	if tx := d.db.Create(user); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (d *ds) GetUser(id string) (*User, error) {
	return nil, nil
}
