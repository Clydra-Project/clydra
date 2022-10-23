//go:build wireinject
// +build wireinject

package command

import (
	"context"

	"clydra.io/services/api"
	"clydra.io/services/config"
	"clydra.io/services/database"
	"clydra.io/services/database/migrator"
	"clydra.io/services/pipeline/pipelineimpl"
	"github.com/google/wire"
)

var wireServerSet = wire.NewSet(
	migrator.NewMigratorService,
	pipelineimpl.NewPipelineService,
	database.NewDatabaseService,
	config.NewConfigService,
	api.NewAPIService,
	NewServer,
)

func InitServer(ctx context.Context) *Server {
	wire.Build(wireServerSet)
	return nil
}
