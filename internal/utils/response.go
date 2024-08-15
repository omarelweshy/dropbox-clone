package util

import (
	"dropbox-clone/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, statusCode int, message string, errors interface{}) {
	c.JSON(statusCode, model.ErrorResponse{
		Status:  "error",
		Message: message,
		Data:    errors,
	})
}

func RespondWithSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, model.APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func RenderLayout(c *gin.Context, content string, data gin.H) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userData, _ := user.(*model.User)
	data["Name"] = userData.Name
	data["Email"] = userData.Email
	data["Avatar"] = userData.Avatar
	data["content"] = content
	c.HTML(http.StatusOK, "layout.templ", data)
}
