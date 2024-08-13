package handler

import (
	"dropbox-clone/internal/form"
	"dropbox-clone/internal/service"
	util "dropbox-clone/internal/utils"
	"fmt"
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
	err := h.FolderService.CreateFolder(form.Name)

	if err != nil {
		fmt.Println(err.Error)
		util.RespondWithError(c, http.StatusInternalServerError, "could not", nil)
		return
	}
	util.RespondWithSuccess(c, "user data", gin.H{
		// "firstName": user.FirstName,
		// "lastName":  user.LastName,
		// "username":  user.Username,
		// "email":     user.Email,
	})
}
