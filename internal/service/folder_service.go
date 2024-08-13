package service

import (
	"dropbox-clone/internal/model"
	"dropbox-clone/internal/repository"
	util "dropbox-clone/internal/utils"
	"errors"
)

type FolderService struct {
	FolderRepository repository.FolderRepository
}

func (s *FolderService) CreateFolder(user_id uint, name string, parentID *string) (*model.Folder, error) {
	folder, err := s.FolderRepository.GetFolderByNameAndParentID(user_id, name, parentID)
	if folder != nil {
		return nil, errors.New("Folder already exists")
	}
	// Return any other error
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// Create a new folder happily and return it if ok or error
	newFolder := model.Folder{
		ID:       util.GenerateUUID(),
		ParentID: parentID,
		Name:     name,
		UserID:   1,
	}
	return s.FolderRepository.CreateFolder(&newFolder)
}
