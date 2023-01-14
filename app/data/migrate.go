package data

func Migrate() error {
	s, err := getDS()
	if err != nil {
		return err
	}
	err = s.db.AutoMigrate(
		&AccessToken{},
		&Follow{},
		&Image{},
		&Post{},
		&PostImage{},
		&User{},
	)
	if err != nil {
		return err
	}
	return nil
}
