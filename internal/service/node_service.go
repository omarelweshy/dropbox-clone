package service

import (
	"dropbox-clone/internal/model"
	"dropbox-clone/internal/repository"
	util "dropbox-clone/internal/utils"
	"errors"
	"strings"
)

type NodeService struct {
	NodeRepository repository.NodeRepository
}

func (s *NodeService) CreateNode(nodeType string, user_id uint, name string, parentID *string) (*model.Node, error) {
	node, err := s.NodeRepository.GetNodeByTypeAndNameAndParentID(nodeType, user_id, name, parentID)
	if node != nil {
		return nil, errors.New(strings.Title(nodeType) + " " + name + " already exists")
	}

	// Return any other error
	if err != nil && err.Error() != "record not found" {
		return nil, errors.New(err.Error())
	}

	// Create a new node happily and return it if ok or error
	newNode := model.Node{
		ID:       util.GenerateUUID(),
		ParentID: parentID,
		Type:     model.NodeType(nodeType),
		Name:     name,
		UserID:   user_id,
	}
	createdNode, err := s.NodeRepository.CreateNode(&newNode)

	if parentID != nil {
		parentNode, err := s.NodeRepository.GetNodeByID(*parentID)
		if err != nil {
			return nil, err
		}
		parentNode.Children = append(parentNode.Children, *createdNode)
	}

	return createdNode, nil
}

func (s *NodeService) ListNode(user_id uint, parentID *string) ([]*model.Node, error) {
	return s.NodeRepository.ListNodes(user_id, parentID)
}
