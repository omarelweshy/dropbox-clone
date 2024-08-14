package handler

import (
	"dropbox-clone/internal/form"
	"dropbox-clone/internal/service"
	util "dropbox-clone/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NodeHandler struct {
	NodeService service.NodeService
}

func (h *NodeHandler) CreateNode(c *gin.Context) {
	var form form.CreateNodeForm

	if err := c.ShouldBindJSON(&form); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := util.FormatValidationError(validationErrors)
			util.RespondWithError(c, http.StatusBadRequest, "Validation failed", formattedErrors)
			return
		}
	}

	node, err := h.NodeService.CreateNode(form.Type, 1, form.Name, form.ParentID)
	if err != nil {
		util.RespondWithError(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.RespondWithSuccess(c, "Node Created", gin.H{
		"ID":   node.ID,
		"Name": node.Name,
	})
}

func (h *NodeHandler) ListNode(c *gin.Context) {
	id := c.Param("id")
	var parentID *string
	if id != "" {
		parentID = &id
	} else {
		parentID = nil
	}
	nodes, _ := h.NodeService.ListNode(1, parentID)

	var nodeResponses []*form.NodeResponse
	for _, node := range nodes {
		nodeResponses = append(nodeResponses, &form.NodeResponse{
			ID:       node.ID,
			ParentID: node.ParentID,
			Type:     string(node.Type),
			Name:     node.Name,
			// Children:  node.Children,
			CreatedAt: node.CreatedAt,
			UpdatedAt: node.UpdatedAt,
		})
	}

	util.RespondWithSuccess(c, "Listing Nodes", gin.H{
		"nodes": nodeResponses,
	})
}
