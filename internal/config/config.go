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
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	SecretKey = os.Getenv("SECRET_KEY")
	DBUrl = os.Getenv("DATABASE_URL")
	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	ClientID = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")

	if ClientID == "" || ClientSecret == "" {
		log.Fatalf("ClientID or ClientSecret is not set")
	}
}
