package handlers

import (
	"educational-service/models"
	"educational-service/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResearchHandler struct {
	service *services.ResearchService
}

func NewResearchHandler(service *services.ResearchService) *ResearchHandler {
	return &ResearchHandler{service: service}
}

func (h *ResearchHandler) GetAllPublications(c *gin.Context) {
	publications, err := h.service.GetAllPublications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, publications)
}

func (h *ResearchHandler) CreatePublication(c *gin.Context) {
	var publication models.Publication
	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreatePublication(&publication); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, publication)
}

func (h *ResearchHandler) CreateConference(c *gin.Context) {
	var conference models.Conference
	if err := c.ShouldBindJSON(&conference); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateConference(&conference); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, conference)
}

func (h *ResearchHandler) GetAllTheses(c *gin.Context) {
	theses, err := h.service.GetAllTheses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, theses)
}

func (h *ResearchHandler) CreateThesis(c *gin.Context) {
	var thesis models.Thesis
	if err := c.ShouldBindJSON(&thesis); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateThesis(&thesis); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, thesis)
}

func (h *ResearchHandler) AssignSupervisor(c *gin.Context) {
	thesisIDStr := c.Param("thesis_id")
	supervisorIDStr := c.Param("supervisor_id")
	thesisID, _ := strconv.ParseUint(thesisIDStr, 10, 64)
	supervisorID, _ := strconv.ParseUint(supervisorIDStr, 10, 64)
	if err := h.service.AssignSupervisor(uint(thesisID), uint(supervisorID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Supervisor assigned"})
}