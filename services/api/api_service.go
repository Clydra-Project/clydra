package api

import (
	"context"

	"clydra.io/services/config"
	"clydra.io/services/pipeline"
	"github.com/gin-gonic/gin"
)

type APIService struct {
	piplineService pipeline.PipelineService
	configService  *config.ConfigService
}

func NewAPIService(piplineService pipeline.PipelineService, configService *config.ConfigService) *APIService {
	return &APIService{
		piplineService: piplineService,
		configService:  configService,
	}
}

func (s *APIService) Run(ctx context.Context) {
	gin.SetMode(s.configService.GetAppMode())
	gin := gin.Default()
	router := s.router(ctx, gin)
	router.Run(s.configService.GetAppPort())
}
