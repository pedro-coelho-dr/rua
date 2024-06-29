package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	UserID    uint   `gorm:"not null"`
	ArtworkID *uint
	ProjectID *uint
	DeletedAt soft_delete.DeletedAt `gorm:"index"`
	CreatedAt time.Time             `gorm:"not null"`
	UpdatedAt time.Time             `gorm:"not null"`
}
