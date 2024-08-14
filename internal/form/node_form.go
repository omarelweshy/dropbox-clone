package form

import (
	"time"
)

type CreateNodeForm struct {
	Name     string  `json:"name,binding:required"`
	Type     string  `json:"Type,binding:required"`
	ParentID *string `json:"parentID,omitempty"`
}

type NodeResponse struct {
	ID        string          `json:"ID"`
	ParentID  *string         `json:"ParentID"`
	Name      string          `json:"Name"`
	Type      string          `json:"Type"`
	Children  []*NodeResponse `json:"Children,omitempty"`
	CreatedAt time.Time       `json:"CreatedAt"`
	UpdatedAt time.Time       `json:"UpdatedAt"`
	FileSize  *int64          `json:"FileSize"`
	FileType  *string         `json:"FileType"`
}
