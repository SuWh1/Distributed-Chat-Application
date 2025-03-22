package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig struct {
	DB struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
		SSLMode  string
	}
	SecretKey string
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env file")
	}

	// DB
	AppConfig.DB.User = os.Getenv("DB_USER")
	AppConfig.DB.Password = os.Getenv("DB_PASSWORD")
	AppConfig.DB.Host = os.Getenv("DB_HOST")
	AppConfig.DB.Port = os.Getenv("DB_PORT")
	AppConfig.DB.Name = os.Getenv("DB_NAME")
	AppConfig.DB.SSLMode = os.Getenv("DB_SSLMODE")

	// Secret key
	AppConfig.SecretKey = os.Getenv("SECRET_KEY")
}
