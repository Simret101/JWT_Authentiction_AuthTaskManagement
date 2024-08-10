package router

import (
	"task/controllers"
	"task/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authorized := r.Group("/")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("/tasks", controllers.GetAllTasks)
		authorized.GET("/tasks/:id", controllers.GetTaskByID)
		authorized.POST("/tasks", controllers.CreateTask)
		authorized.PUT("/tasks/:id", controllers.UpdateTask)
		authorized.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	return r
}
