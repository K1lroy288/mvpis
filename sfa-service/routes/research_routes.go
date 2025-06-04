package routes

import (
	"educational-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupResearchRoutes(router *gin.Engine, handler *handlers.ResearchHandler) {
	api := router.Group("/api/v1/research")
	{
		// Публикации
		api.GET("/publications", handler.GetAllPublications)
		api.POST("/publications", handler.CreatePublication)
		// Конференции
		api.POST("/conferences", handler.CreateConference)
		// Дипломные работы
		api.GET("/theses", handler.GetAllTheses)
		api.POST("/theses", handler.CreateThesis)
		// Привязка руководителя
		api.PUT("/theses/:thesis_id/supervisor/:supervisor_id", handler.AssignSupervisor)
	}
}