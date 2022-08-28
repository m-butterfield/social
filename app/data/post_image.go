package data

type PostImage struct {
	PostID  string `gorm:"primaryKey"`
	Post    *Post
	ImageID string `gorm:"primaryKey"`
	Image   *Image
}
