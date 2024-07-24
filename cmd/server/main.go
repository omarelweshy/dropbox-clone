package main

import (
	"dropbox-clone/internal/model"
	"dropbox-clone/internal/router"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB() (*gorm.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	db, err := setupDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&model.User{})

	r := router.SetupRouter(db)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
