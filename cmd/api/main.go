package main

import (
	"gin-api/internal/config"
	routesHandler "gin-api/internal/handlers"
	usersHandler "gin-api/internal/handlers/users"
	usersRepository "gin-api/internal/repositories/users"
	usersService "gin-api/internal/services/users"
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

	// Initialize repositories
	userRepo := usersRepository.NewUserRepository(dbpool)

	// Initialize services
	userService := usersService.NewUserService(userRepo)

	// Initialize handlers
	userHandler := usersHandler.NewUserHandler(userService)

	// Initialize Gin
	r := gin.Default()

	// Setup routes
	routesHandler.SetupRoutes(r, userHandler)

	// Start the server
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
