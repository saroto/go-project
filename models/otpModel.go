package models

import (
	"time"
)

type Otp struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"` // Use uint for IDs
	UserID    uint      `gorm:"not null;index"`           // Add index for faster lookup
	OtpCode   string    `gorm:"size:6;not null"`          // Limit length to 6 digits
	ExpiresAt time.Time `gorm:"not null;index"`           // Add index for expiry checks
	CreatedAt time.Time
	UpdatedAt time.Time
}
