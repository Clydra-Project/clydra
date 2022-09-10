package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartHttpHandler() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})
	r.Run()
}
