package routes

import (
	"methodological-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler *handlers.DisciplineHandler) {
	api := router.Group("/api/v1/methodology")
	{
		api.GET("/disciplines", handler.GetAllDisciplines)
		api.GET("/disciplines/:id", handler.GetDisciplineByID)
		api.POST("/disciplines", handler.CreateDiscipline)
		api.POST("/disciplines/:id/files", handler.UploadFile)
	}
}
