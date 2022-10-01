package database

import (
	"context"
	"fmt"

	"clydra.io/services/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseService interface {
	GetDB(ctx context.Context) *gorm.DB
	Init(ctx context.Context)
}

type databaseServiceImpl struct {
	db         *gorm.DB
	cfgService *config.ConfigService
}

func (d *databaseServiceImpl) GetDB(ctx context.Context) *gorm.DB {
	return d.db
}

func (d *databaseServiceImpl) Init(ctx context.Context) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", d.cfgService.GetDbHost(), d.cfgService.GetDbUser(), d.cfgService.GetDbPassword(), d.cfgService.GetDbName(), d.cfgService.GetDbPort())
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	d.db = db
}

func NewDatabaseService(ctx context.Context, cfg *config.ConfigService) DatabaseService {
	db := &databaseServiceImpl{
		cfgService: cfg,
	}
	return db
}
