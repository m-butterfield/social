package data

import "time"

type Follow struct {
	FollowerID string    `gorm:"primaryKey;type:uuid;" json:"-"`
	UserID     string    `gorm:"primaryKey;type:uuid;" json:"-"`
	CreatedAt  time.Time `gorm:"not null;default:now()"`
}

func (s *ds) CreateFollow(follow *Follow) error {
	if tx := s.db.Create(follow); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) GetUserFollows(userID string) ([]*Follow, error) {
	var follows []*Follow
	tx := s.db.
		Where("user_id = ?", userID).
		Find(&follows)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return follows, nil
}
