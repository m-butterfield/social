package data

type PostImage struct {
	PostID  int `gorm:"primaryKey"`
	Post    *Post
	ImageID string `gorm:"primaryKey"`
	Image   *Image
}
