package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserId    uint   `json:"userid,omitempty" binding:"required"`
	Name      string `gorm:"unique" json:"path,omitempty" binding:"required"`
	Path      string `gorm:"unique" json:"path,omitempty" binding:"required"`
	Type      string `json:"type,omitempty" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
