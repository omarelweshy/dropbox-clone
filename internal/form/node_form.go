package form

import (
	"dropbox-clone/internal/model"
	"time"
)

type CreateFolderForm struct {
	Name     string  `json:"name,binding:required"`
	ParentID *string `json:"parentID,omitempty"`
}

type NodeResponse struct {
	ID        string       `json:"id"`
	UserID    uint         `json:"UserID"`
	ParentID  *string      `json:"parentID"`
	Name      string       `json:"name"`
	Type      string       `json:"Type"`
	Children  []model.Node `json:"Children"`
	CreatedAt time.Time    `json:"CreatedAt"`
	UpdatedAt time.Time    `json:"UpdatedAt"`
	FileSize  *int64       `json:"FileSize"`
	FileType  *string      `json:"FileType"`
}
