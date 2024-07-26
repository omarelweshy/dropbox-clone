package router

import (
	"dropbox-clone/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/auth/google/login", authHandler.GoogleLogin)
	r.GET("/auth/google/callback", authHandler.GoogleCallback)
	r.POST("/logout", authHandler.Logout)

	return r
}
