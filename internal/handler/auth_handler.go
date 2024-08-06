package handler

import (
	"dropbox-clone/internal/config"
	"dropbox-clone/internal/model"
	"dropbox-clone/internal/service"
	store "dropbox-clone/internal/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	var googleOauthConfig = &oauth2.Config{
		RedirectURL:  config.Host + "/auth/google/callback",
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	url := googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	tokenInfoReq, err := http.Post("https://oauth2.googleapis.com/token"+"?code="+code+"&client_id="+config.ClientID+"&client_secret="+config.ClientSecret+"&grant_type=authorization_code&redirect_uri="+config.Host+"/auth/google/callback", "", nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer tokenInfoReq.Body.Close()
	tokenBody, err := ioutil.ReadAll(tokenInfoReq.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	var response model.GoogleTokenResponse
	if err := json.Unmarshal(tokenBody, &response); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	// ============================================
	accountInfo, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v1/userinfo", nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	accountInfo.Header.Set("Authorization", "Bearer "+response.AccessToken)
	resp, err := http.DefaultClient.Do(accountInfo)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()
	http.Get("https://accounts.google.com/o/oauth2/revoke?token=" + response.AccessToken)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var userInfo model.GoogleUser
	if err := json.Unmarshal(body, &userInfo); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	user := &model.User{
		GoogleID: userInfo.GoogleID,
		Email:    userInfo.Email,
		Name:     userInfo.Name,
		Avatar:   userInfo.Avatar,
	}
	existingUser, err := h.AuthService.AuthRepository.GetByGoogleID(user.GoogleID)

	if err != nil {
		h.AuthService.AuthRepository.CreateUser(user)
	} else {
		user = existingUser
	}
	session, _ := store.GetSession(c.Request)
	session.Values["user_id"] = user.ID
	session.Values["authenticated"] = true

	if err := session.Save(c.Request, c.Writer); err != nil {
		log.Printf("Error saving session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.Redirect(http.StatusFound, "/")
	c.Abort()
	return
}

func (h *AuthHandler) Home(c *gin.Context) {
	session, _ := store.GetSession(c.Request)
	userID, ok := session.Values["user_id"].(uint)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user, err := h.AuthService.AuthRepository.GetByID(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.HTML(http.StatusOK, "layout.templ", gin.H{
		"Title":   "Home",
		"Header":  "Home",
		"content": "index.templ",
		"user":    user,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	session, _ := store.GetSession(c.Request)
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusFound, "/login")
	c.Abort()
	return
}
