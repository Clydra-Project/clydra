package system

import "os"

var config *Config

type Config struct {
	Database string
}

func InitConfig() {
	config = &Config{
		Database: os.Getenv("DB_DSN"),
	}
}

func GetConfig() *Config {
	return config
}
