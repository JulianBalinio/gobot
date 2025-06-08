package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey    string
	SecretKey string
	BaseUrl   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		ApiKey:    os.Getenv("BINGX_API_KEY"),
		SecretKey: os.Getenv("BINGX_SECRET_KEY"),
		BaseUrl:   os.Getenv("BINGX_BASE_URL"),
	}
}
