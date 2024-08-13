package handler

import (
	"dropbox-clone/internal/form"
	"dropbox-clone/internal/service"
	util "dropbox-clone/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FolderHandler struct {
	FolderService service.FolderService
}

func (h *FolderHandler) CreateFolder(c *gin.Context) {
	var form form.CreateFolderForm

	if err := c.ShouldBindJSON(&form); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := util.FormatValidationError(validationErrors)
			util.RespondWithError(c, http.StatusBadRequest, "Validation failed", formattedErrors)
			return
		}
	}

	folder, err := h.FolderService.CreateFolder(1, form.Name, form.ParentID)
	if err != nil {
		util.RespondWithError(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.RespondWithSuccess(c, "Folder Created", gin.H{
		"id":   folder.ID,
		"name": folder.Name,
	})
}
