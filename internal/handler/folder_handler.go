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
	FolderService *service.FolderService
}

func (h FolderHandler) Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "YOU ARE HERE"})
}

func (h FolderHandler) CreateFolder(c *gin.Context) {
	var form form.CreateFolderForm
	if err := c.ShouldBindJSON(&form); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := util.FormatValidationError(validationErrors)
			util.RespondWithError(c, http.StatusBadRequest, "Validation failed", formattedErrors)
			return
		}
	}
	// err := h.FolderService.CreateFolder(&form.Name, *form.ParentID)
	util.RespondWithSuccess(c, "Registration successful", nil)

}
