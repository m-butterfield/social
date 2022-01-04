package data

type Follower struct {
	UserID     string `gorm:"primaryKey"`
	User       *User
	FollowerID string `gorm:"primaryKey"`
	Follower   *User
}

func (s *ds) CreateFollower(string, string) error {
	return nil
}

func (s *ds) DeleteFollower(string, string) error {
	return nil
}
