package repository

import (
	"dropbox-clone/internal/model"
	"errors"

	"gorm.io/gorm"
)

type NodeRepository struct {
	DB *gorm.DB
}

func (r *NodeRepository) CreateNode(node *model.Node) (*model.Node, error) {
	err := r.DB.Create(node).Error
	return node, err
}
func (r *NodeRepository) GetNodeByID(id string) (*model.Node, error) {
	var node model.Node
	if err := r.DB.Where("id = ?", id).First(&node).Error; err != nil {
		return nil, err
	}
	return &node, nil
}

func (r *NodeRepository) UpdateNode(node *model.Node) error {
	return r.DB.Save(node).Error
}

func (r *NodeRepository) ListNodes(userID uint, parentID *string) ([]*model.Node, error) {
	var nodes []*model.Node
	query := r.DB.Where("user_id = ?", userID)
	if parentID == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", parentID)
	}
	if err := query.Find(&nodes).Error; err != nil {
		return nil, err
	}
	if err := query.Preload("Children").Find(&nodes).Error; err != nil {
		return nil, err
	}
	return nodes, nil
}

func (r *NodeRepository) GetNodeByTypeAndNameAndParentID(nodeType string, user_id uint, name string, parentID *string) (*model.Node, error) {
	var node model.Node
	// Check if there and error in the parent id itself
	if parentID != nil {
		var parentNode model.Node
		if err := r.DB.Where("id = ?", parentID).First(&parentNode).Error; err != nil {
			return nil, errors.New("invalid parentID")
		}
	}
	// we split it like that cuz if the parentID is null should be quering like IS NULL
	// But if the parentID is string should be quering like = 'string'
	query := r.DB.Where("user_id = ? AND name = ? AND type = ?", user_id, name, nodeType)
	if parentID == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", parentID)
	}
	if err := query.First(&node).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return &node, nil
}
