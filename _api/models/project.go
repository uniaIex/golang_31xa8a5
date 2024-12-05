package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name,omitempty" binding:"required" gorm:"default:'-'"`
	Description string `json:"description,omitempty" binding:"required" gorm:"default:'-'"`
	Github      string `json:"github,omitempty" binding:"required" gorm:"default:'-'"`
	Deployment  string `json:"deployment,omitempty" binding:"required" gorm:"default:'-'"`
	Technology  string `json:"technology,omitempty" binding:"required" gorm:"default:'-'"`
	Platform    string `json:"platform,omitempty" binding:"required" gorm:"default:'-'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
