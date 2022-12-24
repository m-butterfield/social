package data

type Follow struct {
	FollowerID string `gorm:"type:uuid;not null" json:"-"`
	UserID     string `gorm:"type:uuid;not null" json:"-"`
}

func (s *ds) CreateFollow(follow *Follow) error {
	if tx := s.db.Create(follow); tx.Error != nil {
		return tx.Error
	}
	return nil
}
