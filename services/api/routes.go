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

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			s.routerV1(ctx, v1)
		}
	}
	return router
}

func (s *APIServiceImpl) routerV1(ctx context.Context, v1 *gin.RouterGroup) {

	// pipeline routes
	pipeline := v1.Group("/pipelines")
	{
		pipeline.GET("/", s.getPipelines)
		pipeline.GET("/:id", s.getPipeline)
	}

	// create someting here
	// ...

}
