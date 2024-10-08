package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SecretKey    string
	Host         string
	ClientID     string
	ClientSecret string
	Port         string
	DBUrl        string
	SessionKey   string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

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
