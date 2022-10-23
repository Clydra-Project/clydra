package config

import (
	"context"
	"os"

	"github.com/joho/godotenv"
)

type ConfigService struct {
	// app config
	appPort string
	appMode string

	// database config
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

func (c *ConfigService) LoadFromDotEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	c.LoadConfig()
	return nil
}

func NewConfigService(ctx context.Context) *ConfigService {
	return &ConfigService{}
}
