//go:build wireinject
// +build wireinject

package server

import (
	"context"

	"clydra.io/services/api"
	"clydra.io/services/config"
	"clydra.io/services/database"
	"clydra.io/services/database/migrator"
	"clydra.io/services/pipeline"
	"github.com/google/wire"
)

var wireServerSet = wire.NewSet(
	migrator.NewMigratorService,
	pipeline.NewPipelineService,
	database.NewDatabaseService,
	config.NewConfigFromDotEnv,
	api.NewAPIService,
	NewServer,
)

func InitServer(ctx context.Context) *Server {
	wire.Build(wireServerSet)
	return nil
}
