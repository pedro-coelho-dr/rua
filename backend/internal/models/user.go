package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID          uint                    `gorm:"primaryKey"`
	Name        string                  `gorm:"size:100;not null;index"`
	Email       string                  `gorm:"uniqueIndex;size:100;not null"`
	Password    string                  `gorm:"size:255;not null"`
	AvatarURL   *string                 `gorm:"size:255"`
	Preferences *map[string]interface{} `gorm:"type:jsonb"`
	CreatedAt   time.Time               `gorm:"not null"`
	UpdatedAt   time.Time               `gorm:"not null"`
	DeletedAt   soft_delete.DeletedAt   `gorm:"index"`

	// Relationships Slices
	Roles    []Role    `gorm:"many2many:user_roles;"`
	Photos   []Photo   `gorm:"foreignKey:UserID"`
	Artworks []Artwork `gorm:"foreignKey:UserID"`
	Projects []Project `gorm:"many2many:user_projects;"`
	Comments []Comment `gorm:"foreignKey:UserID"`
	Messages []Message `gorm:"foreignKey:SenderID"`
}
