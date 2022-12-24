package data

import "time"

type Follow struct {
	FollowerID string    `gorm:"type:uuid;not null" json:"-"`
	UserID     string    `gorm:"type:uuid;not null" json:"-"`
	CreatedAt  time.Time `gorm:"not null;default:now()"`
}

func (s *ds) CreateFollow(follow *Follow) error {
	if tx := s.db.Create(follow); tx.Error != nil {
		return tx.Error
	}
	return nil
}
