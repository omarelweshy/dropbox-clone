package middleware

import (
	store "dropbox-clone/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	session, _ := store.GetSession(c.Request)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	c.Next()
}
