package models

import (
	"time"
)

type Post struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Title     string
	Body      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
