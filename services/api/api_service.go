package api

import (
	"context"

	"clydra.io/services/config"
	"clydra.io/services/pipeline"
	"github.com/gin-gonic/gin"
)

type APIService interface {
	Run(ctx context.Context) error
}

type APIServiceImpl struct {
	piplineService pipeline.PipelineService
	configService  *config.ConfigService
}

func NewAPIService(piplineService pipeline.PipelineService, configService *config.ConfigService) APIService {
	return &APIServiceImpl{
		piplineService: piplineService,
		configService:  configService,
	}
}

func (s *APIServiceImpl) Run(ctx context.Context) error {
	gin.SetMode(s.configService.GetAppMode())
	gin := gin.Default()
	router := s.router(ctx, gin)
	return router.Run(s.configService.GetAppPort())
}
