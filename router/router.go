package router

import (
	"task-manager/controller"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// Home page
	r.GET("/", controller.HomePage)
	// Public routes
	r.POST("/register", controller.RegisterUser)
	r.POST("/login", controller.LoginUser)

	// Protected routes
	task := r.Group("/tasks")
	task.Use(middleware.JWTAuth())
	{
		task.POST("", controller.CreateTask)
		task.GET("", controller.GetTasks)
		task.DELETE("/:id", controller.DeleteTask)
	}

	return r
}
