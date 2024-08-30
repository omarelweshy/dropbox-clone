package middleware

import (
	"dropbox-clone/internal/service"
	store "dropbox-clone/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.GetSession(c.Request)
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		userID, ok := session.Values["user_id"].(uint)
		if !ok {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		user, err := authService.AuthRepository.GetByID(userID)
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Set("user_id", user.ID)
		c.Next()
	}
}
