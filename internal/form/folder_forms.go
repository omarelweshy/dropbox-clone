package form

type CreateFolderForm struct {
	Name     string  `json:"name" binding:"required"`
	ParentID *string `json:"parent_id"`
}
