package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	router.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title": "Test",
		})
	})
}
