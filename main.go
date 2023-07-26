package main

import (
	"log"
	"net/http"

	"github.com/NickVTech/simple_api/initializers"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	// Load enviroment variables
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load enviroment variables", err)
	}

	// Connection pool to postgres database
	initializers.ConnectDB(&config)

	// Gin router
	server = gin.Default()
}

func main() {
	// Load enviroment variables
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load enviroment variables", err)
	}

	// Groups routes with common middleware to same path prefix
	router := server.Group("/api")

	// Endpoint, and handler
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})

	})

	router.GET("/users", GetUsers)

	// Run the server on the enviroment port
	log.Fatal(server.Run(":" + config.ServerPort))
}
