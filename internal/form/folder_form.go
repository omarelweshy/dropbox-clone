package form

type CreateFolderForm struct {
	Name     string  `json:"name"`
	ParentID *string `json:"parentID,omitempty"`
}
