package pipeline

type Pipeline struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type PipelineService interface {
	GetPipeline(id int64) (*Pipeline, error)
	GetPipelines() ([]*Pipeline, error)
}

type PipelineRepository interface {
	GetPipeline() (*Pipeline, error)
}
