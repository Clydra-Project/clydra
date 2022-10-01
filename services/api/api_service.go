package api

import "clydra.io/services/pipeline"

type APIService struct {
	piplineService pipeline.PipelineService
}

func NewAPIService(piplineService pipeline.PipelineService) *APIService {
	return &APIService{piplineService: piplineService}
}
