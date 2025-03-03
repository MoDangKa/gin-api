package config

type Config struct {
	ServerAddress string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
}

func LoadConfig() *Config {
	return &Config{
		ServerAddress: ":8080",
		DBHost:        "localhost",
		DBPort:        "5432",
		DBUser:        "postgres",
		DBPassword:    "1234",
		DBName:        "gin_api",
	}
}
