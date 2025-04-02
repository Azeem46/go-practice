package main

import (
	"fmt"
	"go-practice/config"
	"go-practice/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	config.ConnectDB()

	// Get the server port from .env
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default to 8080 if not set
	}

	// Start the server
	fmt.Println("Server running on port", port)
	r := routes.SetupRoutes()
	r.Run(":" + port)
}
