package config

import (
	"github.com/joho/godotenv"
	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2/log"
)

type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT, required" `
	DBHost string `env: "DB_HOST, required"`
	DBName string `env: "DB_NAME, required"`
	DBUser string `env: "DB_USER, required"`
	DBPassword string `env: "DB_PASSWORD, required"`
	DBSslMode string `env: "DB_SSLMODE, required"`


}

func NewEnvConfig() *EnvConfig{
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %e", err)
	}
	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to load variables from .env file: %e",err)
	}
	return config

}