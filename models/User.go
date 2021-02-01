package models

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255; unique" json:"username" binding:"required"`
	Email     string    `gorm:"size:100; unique" json:"email" binding:"required,email"`
	Password  string    `gorm:"size:100;" json:"password" binding:"required"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
