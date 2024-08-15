package router

import (
	"dropbox-clone/internal/handler"
	"dropbox-clone/internal/middleware"
	"dropbox-clone/internal/repository"
	"dropbox-clone/internal/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	authRepository := &repository.UserRepository{DB: db}
	authService := &service.AuthService{AuthRepository: *authRepository}
	authHandler := &handler.AuthHandler{AuthService: *authService}

	nodeRepository := &repository.NodeRepository{DB: db}
	nodeService := &service.NodeService{NodeRepository: *nodeRepository}
	nodeHandler := &handler.NodeHandler{NodeService: *nodeService}

	r := gin.Default()
	r.Static("/static", "./internal/static")
	r.LoadHTMLGlob("internal/templates/*")

	r.Use(cors.Default())
	r.Use(middleware.CORSMiddleware())

	// Auth routers
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.templ", nil)
	})
	r.GET("/auth/google/login", authHandler.GoogleLogin)
	r.GET("/auth/google/callback", authHandler.GoogleCallback)
	r.GET("/auth/logout", authHandler.Logout)

	// Pages
	r.GET("/", middleware.AuthMiddleware(authService), nodeHandler.Home)

	// Node endpoints
	r.POST("/node/create", nodeHandler.CreateNode)
	r.GET("/node", nodeHandler.ListNode)
	r.GET("/node/", nodeHandler.ListNode)
	r.GET("/node/:id", nodeHandler.ListNode)

	return r
}
