package main

import (
	"gin-api/internal/config"
	"gin-api/internal/routes"
	"gin-api/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	cfg := config.LoadConfig()

	dbpool, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	r := gin.Default()

	log.SetOutput(&lumberjack.Logger{
		Filename:   utils.GetLogFilename(),
		MaxSize:    10,
		MaxAge:     14,
		MaxBackups: 3,
		Compress:   true,
	})
	r.Use(gin.LoggerWithWriter(log.Writer()))
	r.Use(config.LimitBodySize(10 * 1024))

	routes.SetupRoutes(r, dbpool)

	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
