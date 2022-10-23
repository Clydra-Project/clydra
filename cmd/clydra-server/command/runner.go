package command

import (
	"context"
	"log"

	"clydra.io/services/api"
	"clydra.io/services/config"
	"clydra.io/services/database"
	"clydra.io/services/database/migrator"
)

type Server struct {
	configService   *config.ConfigService
	apiService      api.APIService
	databaseService database.DatabaseService
	migratorService *migrator.MigratorService
}

func NewServer(config *config.ConfigService, apiService api.APIService, databaseService database.DatabaseService, migratorService *migrator.MigratorService) *Server {
	return &Server{
		configService:   config,
		apiService:      apiService,
		databaseService: databaseService,
		migratorService: migratorService,
	}
}

func (s *Server) Run(ctx context.Context) error {
	err := s.configService.LoadFromDotEnv()
	if err != nil {
		return err
	}

	err = s.databaseService.Init(ctx)
	if err != nil {
		return err
	}
	err = s.migratorService.Migrate(ctx)
	if err != nil {
		return err
	}
	err = s.apiService.Run(ctx)
	if err != nil {
		return err
	}
	return nil
}

func RunServer() int {
	ctx := context.Background()
	srv := InitServer(ctx)
	err := srv.Run(ctx)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}
