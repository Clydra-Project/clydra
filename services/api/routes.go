package api

import (
	"context"

	"github.com/gin-gonic/gin"
)

// router gin
func (s *APIServiceImpl) router(ctx context.Context, router *gin.Engine) *gin.Engine {

	// router health
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api")
	{
		v1.GET("/pipeline", s.getPipeline)
		v1.GET("/pipelines", s.getPipelines)
	}
	return router
}
