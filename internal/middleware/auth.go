package middleware

import (
	"dropbox-clone/internal/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	session, _ := config.Store.Get(c.Request, "sessionid")
	fmt.Println(session)
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}
