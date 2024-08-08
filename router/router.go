package router

import (
	"task/controllers"
	"task/middleware"

	"github.com/gin-gonic/gin"
)

// sets up the router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/tasks", controllers.GetAllTasks)
		auth.GET("/tasks/:id", controllers.GetTaskByID)
		auth.POST("/tasks", controllers.CreateTask)
		auth.PUT("/tasks/:id", controllers.UpdateTask)
		auth.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	// Admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/tasks", controllers.GetAllTasks)
		admin.GET("/tasks/:id", controllers.GetTaskByID)
		admin.POST("/tasks", controllers.CreateTask)
		admin.PUT("/tasks/:id", controllers.UpdateTask)
		admin.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	return r
}
