package task

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	router.GET("/tasks", h.GetTasks)
	router.GET("/tasks/:id", h.GetTask)
	router.POST("/tasks", h.CreateTask)
	router.PATCH("/tasks/:id", h.UpdateTask)
	router.DELETE("/tasks/:id", h.DeleteTask)
}
