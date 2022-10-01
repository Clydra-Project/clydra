package pipeline

import (
	"clydra.io/services/database"
)

type PipelineServiceImpl struct {
}

func NewPipelineService(databaseService database.DatabaseService) PipelineService {
	return &PipelineServiceImpl{}
}

func (s *PipelineServiceImpl) GetPipeline(id int64) (*Pipeline, error) {
	return &Pipeline{
		ID: 1, Name: "Pipeline 1",
	}, nil
}

func (s *PipelineServiceImpl) GetPipelines() ([]*Pipeline, error) {
	return []*Pipeline{
		{ID: 1, Name: "Pipeline 1"},
		{ID: 2, Name: "Pipeline 2"},
	}, nil
}
