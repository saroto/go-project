package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Token       string
	ExpiredDate int64
	UserId      int
	IsRevoke    bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
