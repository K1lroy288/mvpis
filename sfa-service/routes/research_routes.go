package routes

import (
	"educational-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupResearchRoutes(router *gin.Engine, handler *handlers.ResearchHandler) {
	api := router.Group("/api/v1/research")
	{
		// Publications routes
		api.GET("/publications", handler.GetAllPublications)
		api.POST("/publications", handler.CreatePublication)

		// Conferences routes
		api.GET("/conferences", handler.GetAllConferences) // Add this line
		api.POST("/conferences", handler.CreateConference)

		// Theses routes
		api.GET("/theses", handler.GetAllTheses)
		api.POST("/theses", handler.CreateThesis)
		api.PUT("/theses/:thesis_id/supervisor/:supervisor_id", handler.AssignSupervisor)
	}
}
