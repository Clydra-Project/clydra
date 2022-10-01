package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigService struct {
	appPort    string
	appMode    string
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

func (c *ConfigService) GetAppPort() string {
	appPort := ":8080"
	if c.appPort != "" {
		appPort = ":" + c.appPort
	}
	return appPort
}

func (c *ConfigService) GetAppMode() string {
	if c.appMode == "" {
		return "debug"
	}
	return c.appMode
}

func (c *ConfigService) LoadConfig() {
	c.dbHost = os.Getenv("DB_HOST")
	c.dbPort = os.Getenv("DB_PORT")
	c.dbUser = os.Getenv("DB_USER")
	c.dbPassword = os.Getenv("DB_PASSWORD")
	c.dbName = os.Getenv("DB_NAME")
	c.appPort = os.Getenv("APP_PORT")
	c.appMode = os.Getenv("APP_MODE")
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
