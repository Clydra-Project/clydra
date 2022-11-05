package pipelineimpl

import (
	"clydra.io/services/database"
	"clydra.io/services/pipeline"
	"gorm.io/gorm"
)

type pipelineServiceImpl struct {
	pipelineRepository pipeline.PipelineRepository
}

func NewPipelineService(databaseService database.DatabaseService) pipeline.PipelineService {
	return &pipelineServiceImpl{
		pipelineRepository: NewPipelineRepository(databaseService),
	}
}

func (s *pipelineServiceImpl) GetPipeline(id int64) (*pipeline.Pipeline, error) {
	results, err := s.pipelineRepository.GetPipeline()
	if err != nil {
		return nil, err
	}
	return results, err
}

func (s *pipelineServiceImpl) GetPipelines() ([]*pipeline.Pipeline, error) {
	return []*pipeline.Pipeline{
		{ID: 1, Name: "Pipeline 1"},
		{ID: 2, Name: "Pipeline 2"},
	}, nil
}

type pipelineRepositoryImpl struct {
	db *gorm.DB
}

func NewPipelineRepository(databaseService database.DatabaseService) pipeline.PipelineRepository {
	return &pipelineRepositoryImpl{
		db: databaseService.GetDB(),
	}
}

func (r *pipelineRepositoryImpl) GetPipeline() (*pipeline.Pipeline, error) {
	// please select record
	pipe := Pipeline{}
	r.db.First(&pipe)
	return pipe.ToEntity(), nil
}
