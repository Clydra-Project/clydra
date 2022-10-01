package server

import (
	"context"

	"clydra.io/services/api"
	"clydra.io/services/config"
	"clydra.io/services/database"
	"clydra.io/services/database/migrator"
)

type Server struct {
	configService   *config.ConfigService
	apiService      *api.APIService
	databaseService database.DatabaseService
	migratorService *migrator.MigratorService
}

func NewServer(config *config.ConfigService, apiService *api.APIService, databaseService database.DatabaseService, migratorService *migrator.MigratorService) *Server {
	return &Server{
		configService:   config,
		apiService:      apiService,
		databaseService: databaseService,
		migratorService: migratorService,
	}
}

func (s *Server) Run(ctx context.Context) {
	s.configService.LoadFromDotEnv()
	s.databaseService.Init(ctx)
	s.migratorService.Migrate(ctx)
	s.apiService.Run(ctx)
}
