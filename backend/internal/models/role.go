package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Role struct {
	ID        uint                  `gorm:"primaryKey"`
	Name      string                `gorm:"size:100;not null"`
	DeletedAt soft_delete.DeletedAt `gorm:"index"`
	CreatedAt time.Time             `gorm:"not null"`
	UpdatedAt time.Time             `gorm:"not null"`
}
