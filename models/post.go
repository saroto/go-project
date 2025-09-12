package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Title     string
	Body      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
