package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Photographer struct {
	ID        uint                  `gorm:"primaryKey"`
	Name      string                `gorm:"size:100;not null"`
	Bio       string                `gorm:"type:text"`
	AvatarURL string                `gorm:"size:255"`
	UserID    *uint                 `gorm:"index"`
	User      User                  `gorm:"foreignKey:UserID"`
	DeletedAt soft_delete.DeletedAt `gorm:"index"`
	CreatedAt time.Time             `gorm:"not null"`
	UpdatedAt time.Time             `gorm:"not null"`
	Photos    []Photo               `gorm:"foreignKey:PhotographerID"`
}
