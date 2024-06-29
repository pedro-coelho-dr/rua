package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Photo struct {
	ID             uint                  `gorm:"primaryKey"`
	URL            string                `gorm:"size:255;not null"`
	UserID         uint                  `gorm:"not null"`
	ArtworkID      uint                  `gorm:"not null"`
	PhotographerID uint                  `gorm:"not null"`
	Photographer   Photographer          `gorm:"foreignKey:PhotographerID"`
	DeletedAt      soft_delete.DeletedAt `gorm:"index"`
	CreatedAt      time.Time             `gorm:"not null"`
	UpdatedAt      time.Time             `gorm:"not null"`
}
