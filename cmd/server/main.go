package main

import (
	"dropbox-clone/internal/config"
	"dropbox-clone/internal/model"
	"dropbox-clone/internal/router"
	"log"
)

func main() {
	config.LoadConfig()

	config.DB.AutoMigrate(&model.User{})
	r := router.SetupRouter(config.DB)
	if err := r.Run(":" + config.Port); err != nil {
		log.Fatal(err)
	}
}
