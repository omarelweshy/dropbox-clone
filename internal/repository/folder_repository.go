package repository

import (
	"dropbox-clone/internal/model"
	"errors"

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

func (r *FolderRepository) GetFolderByNameAndParentID(user_id uint, name string, parentID *string) (*model.Folder, error) {
	var folder model.Folder
	// Check if there and error in the parent id itself
	if parentID != nil {
		var parentFolder model.Folder
		if err := r.DB.Where("id = ?", parentID).First(&parentFolder).Error; err != nil {
			return nil, errors.New("invalid parentID")
		}
	}
	// we split it like that cuz if the parentID is null should be quering like IS NULL
	// But if the parentID is string should be quering like = 'string'
	query := r.DB.Where("user_id = ? AND name = ?", user_id, name)
	if parentID == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", parentID)
	}
	if err := query.First(&folder).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return &folder, nil
}

func (r *FolderRepository) CreateFolder(folder *model.Folder) (*model.Folder, error) {
	err := r.DB.Create(folder).Error
	return folder, err
}

func (r *FolderRepository) UpdateFolder(folder *model.Folder) error {
	return r.DB.Save(folder).Error
}

func (r *FolderRepository) DeleteFolder(id string) error {
	return r.DB.Delete(&model.Folder{}, id).Error
}
