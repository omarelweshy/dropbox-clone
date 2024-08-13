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

var (
	ErrFolderExists = errors.New("Folder already taken")
)

func (s *FolderService) CreateFolder(name string) error {
	_, err := s.FolderRepository.GetFolderByName(name)
	if err == nil {
		return ErrFolderExists
	}
	folder := model.Folder{
		ID:     util.GenerateUUID(),
		Name:   name,
		UserID: 1,
	}
	return s.FolderRepository.CreateFolder(&folder)
}
