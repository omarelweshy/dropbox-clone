package model

import "time"

type Folder struct {
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ParentID  *string `gorm:"type:uuid;"`
	Parent    *Folder `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
