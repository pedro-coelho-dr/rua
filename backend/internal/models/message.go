package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Message struct {
	ID         uint                  `gorm:"primaryKey"`
	Content    string                `gorm:"type:text;not null"`
	SenderID   uint                  `gorm:"not null"`
	ReceiverID uint                  `gorm:"not null"`
	DeletedAt  soft_delete.DeletedAt `gorm:"index"`
	CreatedAt  time.Time             `gorm:"not null"`
	UpdatedAt  time.Time             `gorm:"not null"`
}
