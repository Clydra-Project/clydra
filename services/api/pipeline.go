package api

import "github.com/gin-gonic/gin"

func (s *APIService) getPipeline(c *gin.Context) {
	pipelines, err := s.piplineService.GetPipeline(1)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, pipelines)
}
