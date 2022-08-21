package data

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID          int
	Body        string
	UserID      int `gorm:"not null"`
	User        *User
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	PublishedAt *time.Time
	PostImages  []*PostImage
}

func (s *ds) CreatePost(post *Post) error {
	if tx := s.db.Create(post); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) PublishPost(id int, images []*Image) error {
	post := &Post{ID: id}
	var postImages []*PostImage
	for _, image := range images {
		postImages = append(postImages, &PostImage{Post: post, Image: image})
	}
	if tx := s.db.Create(&postImages); tx.Error != nil {
		return tx.Error
	}
	now := time.Now().UTC()
	if tx := s.db.Model(&post).Updates(&Post{
		PublishedAt: &now,
	}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) GetPost(id int) (*Post, error) {
	var post *Post
	if tx := s.db.First(&post, id); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return post, nil
}

func (s *ds) GetPosts() ([]*Post, error) {
	var posts []*Post
	tx := s.db.
		Preload("PostImages.Image").
		Where("published_at IS NOT NULL").
		Order("created_at DESC").
		Limit(20).
		Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return posts, nil
}

func (s *ds) GetUserPosts(id int) ([]*Post, error) {
	var posts []*Post
	tx := s.db.
		Preload("PostImages.Image").
		Where("user_id = $1", id).
		Where("published_at IS NOT NULL").
		Order("created_at DESC").
		Limit(20).
		Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return posts, nil
}
