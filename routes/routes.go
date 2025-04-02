package routes

import (
	"go-practice/controllers"
	"go-practice/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	AuthRoutes(r) // Authentication routes

	protected := r.Group("/tasks")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/", controllers.CreateTask)
		protected.GET("/", controllers.GetTasks)
		protected.PUT("/:id", controllers.UpdateTask)
		protected.DELETE("/:id", controllers.DeleteTask)
	}
	protectedAdmin := r.Group("/admin")
protectedAdmin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
{
	protectedAdmin.POST("/create-task", controllers.CreateTask)
	protectedAdmin.DELETE("/delete-user/:id", controllers.DeleteUser)
}

	return r
}
