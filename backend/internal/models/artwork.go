package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Artwork struct {
	ID          uint                  `gorm:"primaryKey"`
	Title       string                `gorm:"size:100;not null"`
	Description string                `gorm:"type:text"`
	ArtistID    uint                  `gorm:"not null"`
	Artist      Artist                `gorm:"foreignKey:ArtistID"`
	Latitude    *float64              `gorm:"type:decimal(10,8);not null"`
	Longitude   *float64              `gorm:"type:decimal(11,8);not null"`
	DeletedAt   soft_delete.DeletedAt `gorm:"index"`
	CreatedAt   time.Time             `gorm:"not null"`
	UpdatedAt   time.Time             `gorm:"not null"`

	Photos   []Photo   `gorm:"foreignKey:ArtworkID"`
	Comments []Comment `gorm:"foreignKey:ArtworkID"`
}
