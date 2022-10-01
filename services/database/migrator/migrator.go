package migrator

import (
	"context"

	"clydra.io/services/database"
	"clydra.io/services/pipeline"
)

type MigratorService struct {
	databaseService database.DatabaseService
}

func NewMigratorService(ctx context.Context, databaseService database.DatabaseService) *MigratorService {
	return &MigratorService{
		databaseService: databaseService,
	}
}

func (s *MigratorService) Migrate(ctx context.Context) {
	s.databaseService.GetDB(ctx).AutoMigrate(
		&pipeline.Pipeline{},
	)
}
