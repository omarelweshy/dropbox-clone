package model

import (
	"time"
)

type File struct {
	ID        string `gorm:"primaryKey"` // Unique identifier
	Name      string `gorm:"not null"`   // File name
	UserID    uint
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Size      int64  // Size of the file in bytes
	FileType  string // Type of the file (e.g., pdf, txt)
	FolderID  string `gorm:"index"` // Foreign key for the folder
	CreatedAt time.Time
	UpdatedAt time.Time
}
