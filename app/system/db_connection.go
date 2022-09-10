package system

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

func InitDbConnection() {
	config := GetConfig().Database
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbConnection = db
}

func GetDbConnection() *gorm.DB {
	return dbConnection
}
