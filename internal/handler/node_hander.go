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
	nodes, err := h.NodeService.ListNode(1, parentID)

	if err != nil {
		util.RespondWithError(c, 404, err.Error(), nil)
		return
	}

	var nodeResponses []*form.NodeResponse
	for _, node := range nodes {
		nodeResponses = append(nodeResponses, &form.NodeResponse{
			ID:        node.ID,
			ParentID:  node.ParentID,
			Type:      string(node.Type),
			Name:      node.Name,
			CreatedAt: util.FormatDateTime(node.CreatedAt),
			UpdatedAt: util.FormatDateTime(node.UpdatedAt),
		})
	}

	util.RespondWithSuccess(c, "Listing Nodes", gin.H{
		"nodes": nodeResponses,
	})
}

func (h *NodeHandler) Home(c *gin.Context) {
	userID, _ := c.Get("user_id")
	userIDUint, _ := userID.(uint)

	nodes, _ := h.NodeService.ListNode(userIDUint, nil)
	var nodeResponses []*form.NodeResponse
	for _, node := range nodes {
		nodeResponses = append(nodeResponses, &form.NodeResponse{
			ID:        node.ID,
			ParentID:  node.ParentID,
			Type:      string(node.Type),
			Name:      node.Name,
			FileSize:  node.FileSize,
			FileType:  node.FileType,
			CreatedAt: util.FormatDateTime(node.CreatedAt),
			UpdatedAt: util.FormatDateTime(node.UpdatedAt),
		})
	}

	util.RenderLayout(c, "index.templ", gin.H{
		"Title":  "Home",
		"Header": "All Files",
		"Nodes":  nodeResponses,
	})
}
