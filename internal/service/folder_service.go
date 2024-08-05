package service

import (
	"dropbox-clone/internal/model"
	"dropbox-clone/internal/repository"
	util "dropbox-clone/internal/utils"
)

type FolderService struct {
	Repo repository.FolderRepository
}

func (s *FolderService) CreateFolder(name string, parentID string) (*model.Folder, error) {
	folder := &model.Folder{
		ID:       util.GenerateUUID(),
		Name:     name,
		ParentID: &parentID,
	}
	if err := s.Repo.CreateFolder(folder); err != nil {
		return nil, err
	}
	return folder, nil
}
