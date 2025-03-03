package main

import (
	"gin-api/internal/config"
	"gin-api/internal/handlers"
	"gin-api/internal/repositories"
	"gin-api/internal/services"
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

	// Initialize repository, service, and handler
	userRepo := repositories.NewUserRepository(dbpool)
	userService := services.NewUserService(userRepo)
	handler := handlers.NewHandler(userService)

	// Initialize Gin
	r := gin.Default()

	// Setup routes
	handler.SetupRoutes(r)

	// Start the server
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
