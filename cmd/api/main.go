package main

import (
	"gin-api/internal/config"
	"gin-api/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to PostgreSQL
	dbpool, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	// Initialize Gin
	r := gin.Default()

	// Set up routes with dbpool
	routes.SetupRoutes(r, dbpool)

	// Start the server
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
