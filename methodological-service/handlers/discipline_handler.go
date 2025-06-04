package handlers

import (
	"methodological-service/models"
	"methodological-service/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DisciplineHandler struct {
	service *services.DisciplineService
}

func NewDisciplineHandler(service *services.DisciplineService) *DisciplineHandler {
	return &DisciplineHandler{service: service}
}

func (h *DisciplineHandler) GetAllDisciplines(c *gin.Context) {
	disciplines, err := h.service.GetAllDisciplines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, disciplines)
}

func (h *DisciplineHandler) GetDisciplineByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	discipline, err := h.service.GetDisciplineByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Discipline not found"})
		return
	}
	c.JSON(http.StatusOK, discipline)
}

func (h *DisciplineHandler) CreateDiscipline(c *gin.Context) {
	var discipline models.Discipline
	if err := c.ShouldBindJSON(&discipline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RegisterDiscipline(&discipline); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, discipline)
}

func (h *DisciplineHandler) UploadFile(c *gin.Context) {
	var file models.File
	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UploadFile(&file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, file)
}
