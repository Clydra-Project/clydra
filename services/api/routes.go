package api

import (
	"context"

	"github.com/gin-gonic/gin"
)

// router gin

func (s *APIService) router(ctx context.Context, router *gin.Engine) *gin.Engine {

	// router health
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api")
	{
		v1.GET("/pipeline", s.GetPipeline)
	}
	return router
}

func (s *APIService) Run(ctx context.Context) {
	gin := gin.Default()
	router := s.router(ctx, gin)
	router.Run(":8080")
}
