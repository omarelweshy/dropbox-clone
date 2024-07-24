package handler

import (
	"context"
	"dropbox-clone/internal/model"
	"dropbox-clone/internal/repository"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2Service "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var (
	store             = sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("HOST") + "/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		// Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
)

type AuthHandler struct {
	userRepository repository.UserRepository
}

func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
func (h *AuthHandler) GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	session, _ := store.Get(c.Request, "sessionid")
	if state != session.Values["state"] {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	oauth2Service, err := oauth2Service.NewService(context.Background(), option.WithTokenSource(oauth2.StaticTokenSource(token)))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	userinfoService := oauth2Service.Userinfo.V2.Me
	userinfo, err := userinfoService.Get().Do()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user := &model.User{
		GoogleID: userinfo.Id,
		Email:    userinfo.Email,
		Name:     userinfo.Name,
		Avatar:   userinfo.Picture,
	}
	existingUser, err := h.userRepository.GetByGoogleID(user.GoogleID)
	if err != nil {
		h.userRepository.CreateUser(user)
	} else {
		user = existingUser
	}
	// session.Values["user_id"] = user.ID
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{"message": "User authenticated", "user": user})

}

func (h *AuthHandler) Logout(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessionid")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
