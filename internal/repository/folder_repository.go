package repository

import (
	"dropbox-clone/internal/model"

	"gorm.io/gorm"
)

type FolderRepository struct {
	DB *gorm.DB
}

func (r *FolderRepository) CreateFolder(folder *model.Folder) error {
	return r.DB.Create(folder).Error
}
