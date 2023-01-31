package data

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Body        string `json:"body"`
	Film        string `json:"film"`
	Camera      string `json:"camera"`
	Lens        string `json:"lens"`
	UserID      string `gorm:"not null" json:"-"`
	User        *User
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	PublishedAt *time.Time
	PostImages  []*PostImage
	Images      []*Image `gorm:"many2many:post_images;"`

	CommentCount int `gorm:"-:migration"`
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
	return s.getPosts(nil, before)
}

func (s *ds) GetUserPosts(id string, before *time.Time) ([]*Post, error) {
	return s.getPosts([]string{id}, before)
}

func (s *ds) GetUsersPosts(userIDs []string, before *time.Time) ([]*Post, error) {
	return s.getPosts(userIDs, before)
}

func (s *ds) getPosts(userIDs []string, before *time.Time) ([]*Post, error) {
	var posts []*Post
	if before == nil {
		now := time.Now()
		before = &now
	}
	tx := s.db.
		Select("posts.*, COUNT(comments.post_id) as comment_count").
		Preload("Images").
		Preload("User").
		Where("posts.published_at IS NOT NULL").
		Where("posts.published_at < ?", before).
		Order("posts.created_at DESC").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Limit(2)
	if userIDs != nil {
		tx = tx.Where("posts.user_id in ?", userIDs)
	}
	tx = tx.Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return posts, nil
}
