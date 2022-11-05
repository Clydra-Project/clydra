package database

import (
	"context"
	"fmt"

	"clydra.io/services/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseService interface {
	GetDB() *gorm.DB
	Init(ctx context.Context) error
}

type databaseServiceImpl struct {
	db         *gorm.DB
	cfgService *config.ConfigService
}

func (d *databaseServiceImpl) GetDB() *gorm.DB {
	return d.db
}

func (d *databaseServiceImpl) Init(ctx context.Context) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", d.cfgService.GetDbHost(), d.cfgService.GetDbUser(), d.cfgService.GetDbPassword(), d.cfgService.GetDbName(), d.cfgService.GetDbPort())
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func NewDatabaseService(ctx context.Context, cfg *config.ConfigService) DatabaseService {
	cfg.LoadFromDotEnv()
	db := &databaseServiceImpl{
		cfgService: cfg,
	}
	db.Init(ctx)
	return db
}
