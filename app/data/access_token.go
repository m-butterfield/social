package data

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"gorm.io/gorm"
	"time"
)

const (
	tokenByteLength = 32
	TokenTTL        = 720 * time.Hour
)

type AccessToken struct {
	ID        string
	ExpiresAt time.Time `gorm:"not null"`
	UserID    string    `gorm:"not null"`
	User      *User
}

func (s *ds) CreateAccessToken(user *User) (*AccessToken, error) {
	tokenStr, err := randomToken()
	if err != nil {
		return nil, err
	}
	token := &AccessToken{
		ID:        tokenStr,
		ExpiresAt: time.Now().UTC().Add(TokenTTL),
		User:      user,
	}
	if tx := s.db.Create(token); tx.Error != nil {
		return nil, tx.Error
	}
	return token, nil
}

func (s *ds) DeleteAccessToken(id string) error {
	if tx := s.db.Delete(&AccessToken{}, "id = $1", id); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func randomToken() (string, error) {
	b := make([]byte, tokenByteLength)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func (s *ds) GetAccessToken(id string) (*AccessToken, error) {
	token := &AccessToken{}
	tx := s.db.Preload("User").First(&token, "id = $1", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	if token.ExpiresAt.Before(time.Now().UTC()) {
		if tx := s.db.Delete(token); tx.Error != nil {
			return nil, tx.Error
		}
		return nil, nil
	}
	return token, nil
}
