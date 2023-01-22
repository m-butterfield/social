package data

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Body        string `json:"body"`
	UserID      string `gorm:"not null" json:"-"`
	User        *User
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	PublishedAt *time.Time
	PostImages  []*PostImage
	Images      []*Image `gorm:"many2many:post_images;"`
}

func (s *ds) CreatePost(post *Post) error {
	if tx := s.db.Create(post); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) PublishPost(id string, images []*Image) error {
	var postImages []*PostImage
	for _, image := range images {
		postImages = append(postImages, &PostImage{PostID: id, Image: image})
	}
	if tx := s.db.Create(&postImages); tx.Error != nil {
		return tx.Error
	}
	now := time.Now().UTC()
	if tx := s.db.Model(&Post{ID: id}).Updates(&Post{
		PublishedAt: &now,
	}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) GetPost(id string) (*Post, error) {
	var post *Post
	if tx := s.db.First(&post, "id = $1", id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return post, nil
}

func (s *ds) GetPosts(before *time.Time) ([]*Post, error) {
	var posts []*Post
	if before == nil {
		now := time.Now()
		before = &now
	}
	tx := s.db.
		Preload("Images").
		Preload("User").
		Where("published_at IS NOT NULL").
		Where("published_at < ?", before).
		Order("created_at DESC").
		Limit(2).
		Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return posts, nil
}

func (s *ds) GetUserPosts(id string, before *time.Time) ([]*Post, error) {
	return s.GetUsersPosts([]string{id}, before)
}

func (s *ds) GetUsersPosts(userIDs []string, before *time.Time) ([]*Post, error) {
	var posts []*Post
	if before == nil {
		now := time.Now()
		before = &now
	}
	tx := s.db.
		Preload("Images").
		Preload("User").
		Where("user_id in ?", userIDs).
		Where("published_at IS NOT NULL").
		Where("published_at < ?", before).
		Order("created_at DESC").
		Limit(2).
		Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return posts, nil
}
