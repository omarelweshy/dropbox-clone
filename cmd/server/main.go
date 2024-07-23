package main

import (
	"dropbox-clone/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")
	r.Static("/static", "./web/static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/login-url", func(c *gin.Context) {
		url := auth.GetGoogleLoginURL()
		c.JSON(http.StatusOK, gin.H{"url": url})
	})
	r.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		token, err := auth.ExchangeCodeForToken(code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
			return
		}
		userInfo, err := auth.GetUserInfo(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
			return
		}
		userInfo["token"] = token.AccessToken
		c.JSON(http.StatusOK, userInfo)
	})
	r.Run(":8080")
}
