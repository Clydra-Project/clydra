package pipeline

import (
	"clydra.io/services/database"
)

type pipelineServiceImpl struct {
}

func NewPipelineService(databaseService database.DatabaseService) PipelineService {
	return &pipelineServiceImpl{}
}

func (s *pipelineServiceImpl) GetPipeline(id int64) (*Pipeline, error) {
	return &Pipeline{
		ID: 1, Name: "Pipeline 1",
	}, nil
}

func (s *pipelineServiceImpl) GetPipelines() ([]*Pipeline, error) {
	return []*Pipeline{
		{ID: 1, Name: "Pipeline 1"},
		{ID: 2, Name: "Pipeline 2"},
	}, nil
}
