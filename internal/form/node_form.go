package form

type CreateNodeForm struct {
	Name     string  `json:"Name,binding:required"`
	Type     string  `json:"Type,binding:required"`
	ParentID *string `json:"ParentID,omitempty"`
}

type NodeResponse struct {
	ID        string          `json:"ID"`
	ParentID  *string         `json:"ParentID"`
	Name      string          `json:"Name"`
	Type      string          `json:"Type"`
	Children  []*NodeResponse `json:"Children,omitempty"`
	CreatedAt string          `json:"CreatedAt"`
	UpdatedAt string          `json:"UpdatedAt"`
	FileSize  int64           `json:"FileSize"`
	FileType  string          `json:"FileType"`
}
