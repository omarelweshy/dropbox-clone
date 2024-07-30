package router

import (
	"dropbox-clone/internal/handler"
	"dropbox-clone/internal/middleware"
	"dropbox-clone/internal/repository"
	"dropbox-clone/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	authRepository := &repository.UserRepository{DB: db}
	authService := &service.AuthService{AuthRepository: *authRepository}
	authHandler := &handler.AuthHandler{AuthService: *authService}
	folderHandler := &handler.FolderHandler{}

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.CORSMiddleware())

	// Auth routers
	r.GET("/auth/me", middleware.AuthMiddleware, authHandler.Me)
	r.GET("/auth/google/login", authHandler.GoogleLogin)
	r.GET("/auth/google/callback", authHandler.GoogleCallback)
	r.GET("/auth/logout", authHandler.Logout)

	// Folders routers
	r.GET("/", middleware.AuthMiddleware, folderHandler.Dashboard)
	return r
}
