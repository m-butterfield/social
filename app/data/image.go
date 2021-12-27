package data

import (
	"errors"
	"gorm.io/gorm"
)

type Image struct {
	ID     string `gorm:"type:varchar(128)"`
	Width  int    `gorm:"not null"`
	Height int    `gorm:"not null"`
}

func (s *ds) GetOrCreateImage(id string, width, height int) (*Image, error) {
	image, err := s.getImage(id)
	if err != nil {
		return nil, err
	}
	if image != nil {
		return image, err
	}
	image = &Image{
		ID:     id,
		Width:  width,
		Height: height,
	}
	if tx := s.db.Create(image); tx.Error != nil {
		return nil, tx.Error
	}
	return image, nil
}

func (s *ds) getImage(id string) (*Image, error) {
	image := &Image{}
	tx := s.db.First(&image, "id = $1", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return image, nil
}
