package main

import (
	"dropbox-clone/internal/config"
	"dropbox-clone/internal/model"
	"dropbox-clone/internal/router"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	config.LoadConfig()

	db, err := setupDB()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&model.User{})
	r := router.SetupRouter(db)
	if err := r.Run(":" + config.Port); err != nil {
		log.Fatal(err)
	}
}
