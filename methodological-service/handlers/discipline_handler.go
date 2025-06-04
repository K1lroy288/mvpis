package handlers

import (
	"io"
	"methodological-service/models"
	"methodological-service/services"
	"net/http"
	"os"
	"path/filepath"
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
	idStr := c.Param("id")
	disciplineID, _ := strconv.ParseUint(idStr, 10, 64)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Base(file.Filename)
	uploadPath := filepath.Join("uploads", filename)

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(uploadPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer dst.Close()

	// Copy the content from the source file to the destination file
	_, err = io.Copy(dst, src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newFile := models.File{
		Name:         filename,
		Path:         uploadPath,
		DisciplineID: uint(disciplineID),
	}
	if err := h.service.UploadFile(&newFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "File uploaded successfully", "filename": file.Filename})
}

func (h *DisciplineHandler) GetFilesForDiscipline(c *gin.Context) {
	idStr := c.Param("id")
	disciplineID, _ := strconv.ParseUint(idStr, 10, 64)

	files, err := h.service.GetFilesForDiscipline(uint(disciplineID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, files)
}
