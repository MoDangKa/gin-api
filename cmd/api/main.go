package main

import (
	"context"
	"gin-api/internal/config"
	"gin-api/internal/routes"
	"gin-api/pkg/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	// Load environment variables from .env file (for development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}
}

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	dbpool, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	// Set up logging
	log.SetOutput(&lumberjack.Logger{
		Filename:   utils.GetLogFilename(),
		MaxSize:    10,   // Max size in megabytes
		MaxAge:     14,   // Max age in days
		MaxBackups: 3,    // Max number of old log files to keep
		Compress:   true, // Compress old log files
	})

	// Initialize Gin router
	r := gin.Default()

	// Middleware
	r.Use(gin.LoggerWithWriter(log.Writer()))
	r.Use(config.LimitBodySize(10 * 1024)) // Limit request body size to 10KB

	// Set up routes
	routes.SetupRoutes(r, dbpool)

	// Start server with graceful shutdown
	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
