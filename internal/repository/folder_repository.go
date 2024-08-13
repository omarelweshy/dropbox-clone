package repository

import (
	"dropbox-clone/internal/model"
	"fmt"

	"gorm.io/gorm"
)

type FolderRepository struct {
	DB *gorm.DB
}

func (r *FolderRepository) GetFolderByID(id string) (*model.Folder, error) {
	var folder model.Folder
	if err := r.DB.First(&folder, id).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}

func (r *FolderRepository) GetFolderByName(name string) (*model.Folder, error) {
	var folder model.Folder
	if err := r.DB.First(&folder, name).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}

func (r *FolderRepository) CreateFolder(folder *model.Folder) error {
	fmt.Println("omar")
	return r.DB.Create(folder).Error
}

func (r *FolderRepository) UpdateFolder(folder *model.Folder) error {
	return r.DB.Save(folder).Error
}

func (r *FolderRepository) DeleteFolder(id string) error {
	return r.DB.Delete(&model.Folder{}, id).Error
}
