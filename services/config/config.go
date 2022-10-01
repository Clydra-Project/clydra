package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigService struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
}

func (c *ConfigService) GetDbHost() string {
	return c.dbHost
}

func (c *ConfigService) GetDbPort() string {
	return c.dbPort
}

func (c *ConfigService) GetDbUser() string {
	return c.dbUser
}

func (c *ConfigService) GetDbPassword() string {
	return c.dbPassword
}

func (c *ConfigService) GetDbName() string {
	return c.dbName
}

func (c *ConfigService) LoadConfig() {
	c.dbHost = os.Getenv("DB_HOST")
	c.dbPort = os.Getenv("DB_PORT")
	c.dbUser = os.Getenv("DB_USER")
	c.dbPassword = os.Getenv("DB_PASSWORD")
	c.dbName = os.Getenv("DB_NAME")
}

func (c *ConfigService) LoadFromDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c.LoadConfig()
}

func NewConfigService(ctx context.Context) *ConfigService {
	return &ConfigService{}
}

func NewConfigFromDotEnv(ctx context.Context) *ConfigService {
	cfg := NewConfigService(ctx)
	return cfg
}
