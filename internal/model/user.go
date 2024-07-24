package model

import "time"

type User struct {
	UserID    uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
}
