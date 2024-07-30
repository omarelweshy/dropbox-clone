package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FolderHandler struct{}

func (h FolderHandler) Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "YOU ARE HERE"})
}
