package model

import (
	"time"
)

type Folder struct {
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	UserID    uint
	ParentID  *string  `gorm:"index"` // Self-referencing foreign key for parent folder
	User      User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Files     []File   `gorm:"constraint:OnDelete:CASCADE;"`                     // One-to-many relationship with files
	Folders   []Folder `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE;"` // Recursive relationship
	CreatedAt time.Time
	UpdatedAt time.Time
}
