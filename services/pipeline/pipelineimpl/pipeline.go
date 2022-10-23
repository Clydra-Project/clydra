package pipelineimpl

import (
	"clydra.io/services/database"
	"clydra.io/services/pipeline"
)

type pipelineServiceImpl struct {
}

func NewPipelineService(databaseService database.DatabaseService) pipeline.PipelineService {
	return &pipelineServiceImpl{}
}

func (s *pipelineServiceImpl) GetPipeline(id int64) (*pipeline.Pipeline, error) {
	return &pipeline.Pipeline{
		ID: 1, Name: "Pipeline 1",
	}, nil
}

func (s *pipelineServiceImpl) GetPipelines() ([]*pipeline.Pipeline, error) {
	return []*pipeline.Pipeline{
		{ID: 1, Name: "Pipeline 1"},
		{ID: 2, Name: "Pipeline 2"},
	}, nil
}
