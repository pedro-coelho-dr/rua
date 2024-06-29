package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type Project struct {
	ID          uint                  `gorm:"primaryKey"`
	Title       string                `gorm:"size:100;not null"`
	Description string                `gorm:"type:text"`
	Status      string                `gorm:"size:50;not null;default:'active'"`
	StartDate   *time.Time            `gorm:"index"`
	EndDate     *time.Time            `gorm:"index"`
	FundingGoal *float64              `gorm:"type:decimal(10,2)"`
	FundsRaised *float64              `gorm:"type:decimal(10,2)"`
	DeletedAt   soft_delete.DeletedAt `gorm:"index"`
	CreatedAt   time.Time             `gorm:"not null"`
	UpdatedAt   time.Time             `gorm:"not null"`
	Users       []User                `gorm:"many2many:user_projects"`
	Comments    []Comment             `gorm:"foreignKey:ProjectID"`
	Photos      []Photo               `gorm:"foreignKey:ProjectID"`
}
