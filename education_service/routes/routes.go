package routes

import (
	"educational-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *handlers.StudentHandler) {
	api := router.Group("/api/v1/students")
	{
		api.GET("/", handler.GetAllStudents)
		api.GET("/:id", handler.GetStudentByID)
		api.POST("/", handler.CreateStudent)
		api.POST("/:id/grades", handler.AddGrade)
		api.GET("/honor", handler.GetHonorStudents)       // Добавляем новый endpoint
		api.GET("/expelled", handler.GetExpelledStudents) // Добавляем новый endpoint
	}
}
