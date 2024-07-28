package model

import (
	"time"
)

type User struct {
	UserID    uint   `gorm:"primaryKey"`
	GoogleID  string `gorm:"uniqueIndex"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Avatar    string `json:"avatar"`
	CreatedAt time.Time
}

type GoogleTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type GoogleUser struct {
	GoogleID string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Avatar   string `json:"picture"`
}
