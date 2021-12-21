package data

type Post struct {
	ID     int
	Body   string
	UserID string `gorm:"type:varchar(64);not null"`
	User   *User
}

func (s *ds) CreatePost(post *Post) error {
	if tx := s.db.Create(post); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) GetPosts() ([]*Post, error) {
	var posts []*Post
	if tx := s.db.Find(&posts); tx.Error != nil {
		return nil, tx.Error
	}
	return posts, nil
}
