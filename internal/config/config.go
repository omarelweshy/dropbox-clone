package config

import (
	"log"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	SecretKey    string
	Host         string
	ClientID     string
	ClientSecret string
	Port         string
	DBUrl        string
	DB           *gorm.DB
	Store        *pgstore.PGStore
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

	DB, err = gorm.Open(postgres.Open(DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to session store: %v", err)
	}

	Store, err = pgstore.NewPGStore(DBUrl, []byte(SecretKey))
	if err != nil {
		log.Fatalf("Failed to connect to session store: %v", err)
	}
	defer Store.Close()

	if ClientID == "" || ClientSecret == "" {
		log.Fatalf("ClientID or ClientSecret is not set")
	}
}
