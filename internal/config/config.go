package config

import (
	"log"
	"os"
)

type Config struct {
	ServerAddress string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
}

func LoadConfig() *Config {
	serverAddress, ok := os.LookupEnv("SERVER_ADDRESS")
	if !ok {
		log.Fatal("SERVER_ADDRESS is required but not set")
	}

	DBHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		DBHost = "localhost"
	}

	DBPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		DBPort = "5432"
	}

	DBUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		log.Fatal("DB_USER is required but not set")
	}

	DBPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		log.Fatal("DB_PASSWORD is required but not set")
	}

	DBName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		log.Fatal("DB_NAME is required but not set")
	}

	return &Config{
		ServerAddress: serverAddress,
		DBHost:        DBHost,
		DBPort:        DBPort,
		DBUser:        DBUser,
		DBPassword:    DBPassword,
		DBName:        DBName,
	}
}
