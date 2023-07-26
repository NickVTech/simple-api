package main

import (
	"fmt"
	"log"

	"github.com/NickVTech/simple_api/initializers"
	"github.com/NickVTech/simple_api/models"
)

// Loads config to connect to db
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load enviroment variables", err)
	}

	initializers.ConnectDB(&config)
}

// Migrates data
func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
