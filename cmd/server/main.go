package main

import (
	"dropbox-clone/internal/config"
	"dropbox-clone/internal/model"
	"dropbox-clone/internal/router"
	store "dropbox-clone/internal/utils"
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

	store.InitializeStore([]byte(config.SessionKey))

	if err != nil {
		log.Fatalf("Failed to create PGStore: %v", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Folder{})
	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}

	r := router.SetupRouter(db)
	if err := r.Run(":" + config.Port); err != nil {
		log.Fatal(err)
	}
}
