package models

import (
	"time"
)

type ShortUrl struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Code      string    `json:"code" binding:"required"`
	Link      string    `json:"link" binding:"required,email"`
	Click     uint64    `gorm:"default:0" json:"click"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
