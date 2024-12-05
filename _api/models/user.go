package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"unique" json:"username,omitempty" binding:"required"`
	Email     string `gorm:"unique" json:"email,omitempty" binding:"required"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role,omitempty" binding:"required" gorm:"default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
