package routes

import (
	"go-template/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/tasks", controllers.CreateTaskHandler)
	router.GET("/tasks", controllers.GetTasksHandler)
	router.PUT("/tasks/:id/complete", controllers.CompleteTaskHandler)
	router.GET("/tasks/:id", controllers.GetTaskByIDHandler)
	router.DELETE("/tasks/:id", controllers.DeleteTaskHandler)
	router.DELETE("/tasks/completed", controllers.DeleteCompletedTasksHandler)

}
