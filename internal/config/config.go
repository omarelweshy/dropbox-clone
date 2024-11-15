package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	GinMode      string
	SecretKey    string
	Host         string
	ClientID     string
	ClientSecret string
	Port         string
	DBUrl        string
	SessionKey   string
)

func LoadConfig() {
	dir, errDir := os.Getwd()
	if errDir != nil {
		log.Fatal(errDir)
	}
	environmentPath := filepath.Join(dir, ".env")

	err := godotenv.Load(environmentPath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	GinMode = os.Getenv("ENV_MODE")
	DBUrl = os.Getenv("DATABASE_URL")
	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	ClientID = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")

	SessionKey = os.Getenv("SECRET_KEY")

	if ClientID == "" || ClientSecret == "" {
		log.Fatalf("ClientID or ClientSecret is not set")
	}
}
