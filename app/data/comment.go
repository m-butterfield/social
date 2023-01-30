package data

import (
	"time"
)

type Comment struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Body      string `gorm:"not null"`
	UserID    string `gorm:"not null" json:"-"`
	User      *User
	PostID    string `gorm:"not null" json:"-"`
	Post      *Post
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

func (s *ds) CreateComment(comment *Comment) error {
	if tx := s.db.Create(comment); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) GetComments(postID string, before *time.Time) ([]*Comment, error) {
	var comments []*Comment
	if before == nil {
		now := time.Now()
		before = &now
	}
	tx := s.db.
		Preload("User").
		Where("post_id = ?", postID).
		Where("created_at < ?", before).
		Order("created_at").
		Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return comments, nil
}
