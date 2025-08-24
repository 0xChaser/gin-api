package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Album struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title  string    `json:"title" gorm:"not null"`
	Artist string    `json:"artist" gorm:"not null"`
	Price  float64   `json:"price" gorm:"not null"`
}

func (album *Album) BeforeCreate(tx *gorm.DB) error {
	if album.ID == uuid.Nil {
		album.ID = uuid.New()
	}
	return nil
}
