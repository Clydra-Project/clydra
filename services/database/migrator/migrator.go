package migrator

import (
	"context"

	"clydra.io/services/database"
	"clydra.io/services/pipeline/pipelineimpl"
)

type MigratorService struct {
	databaseService database.DatabaseService
}

func NewMigratorService(ctx context.Context, databaseService database.DatabaseService) *MigratorService {
	return &MigratorService{
		databaseService: databaseService,
	}
}

func (s *MigratorService) Migrate(ctx context.Context) error {
	return s.databaseService.GetDB(ctx).AutoMigrate(

		// pipeline modules
		&pipelineimpl.Pipeline{},
		&pipelineimpl.PipelineNodeType{},
		&pipelineimpl.PipelineNode{},
		&pipelineimpl.PipelineRunner{},
		&pipelineimpl.PipelineLog{},
		&pipelineimpl.PipelineEdgeNode{},
	)
}
