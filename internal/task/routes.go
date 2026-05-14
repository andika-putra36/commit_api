package task

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	router.GET("/tasks", h.GetTasks)
}
