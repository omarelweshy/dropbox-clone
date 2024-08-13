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

func (h *NodeHandler) CreateFolder(c *gin.Context) {
	var form form.CreateFolderForm

	if err := c.ShouldBindJSON(&form); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := util.FormatValidationError(validationErrors)
			util.RespondWithError(c, http.StatusBadRequest, "Validation failed", formattedErrors)
			return
		}
	}

	folder, err := h.NodeService.CreateNode("folder", 1, form.Name, form.ParentID)
	if err != nil {
		util.RespondWithError(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.RespondWithSuccess(c, "Folder Created", gin.H{
		"id":   folder.ID,
		"name": folder.Name,
	})
}

func (h *NodeHandler) ListFolder(c *gin.Context) {
	id := c.Param("id")
	var parentID *string
	if id != "" {
		parentID = &id
	} else {
		parentID = nil
	}
	folders, _ := h.NodeService.ListNode("folder", 1, parentID)

	var folderResponses []*form.NodeResponse
	for _, folder := range folders {
		folderResponses = append(folderResponses, &form.NodeResponse{
			ID:        folder.ID,
			UserID:    folder.UserID,
			ParentID:  folder.ParentID,
			Type:      "folder",
			Name:      folder.Name,
			CreatedAt: folder.CreatedAt,
			UpdatedAt: folder.UpdatedAt,
		})
	}

	util.RespondWithSuccess(c, "Listing Folders", gin.H{
		"folders": folderResponses,
	})
}
