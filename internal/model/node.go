package model

import "time"

type NodeType string

const (
	FolderType NodeType = "folder"
	FileType   NodeType = "file"
)

type Node struct {
	ID        string `gorm:"primaryKey;type:uuid"`
	UserID    uint
	ParentID  *string  `gorm:"type:uuid"`
	Name      string   `gorm:"not null"`
	Type      NodeType `gorm:"not null;default:file"`
	Parent    *Node    `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Children  []Node   `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Additional fields for files
	FileSize int64
	FileType string
}
