package auth

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

func GetGoogleLoginURL() string {
	return googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func ExchangeCodeForToken(code string) (*oauth2.Token, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Could not exchange code for token: %v", err)
		return nil, err
	}
	return token, nil
}

func GetUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
