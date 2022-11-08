package api

import "github.com/gin-gonic/gin"

func (s *APIServiceImpl) getPipeline(c *gin.Context) {
	pipelines, err := s.piplineService.GetPipeline(1)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, pipelines)
}

func (s *APIServiceImpl) getPipelines(c *gin.Context) {
	pipelines, err := s.piplineService.GetPipelines()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, pipelines)
}
