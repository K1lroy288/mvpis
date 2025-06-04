package repositories

import (
	"educational-service/models"

	"gorm.io/gorm"
)

type ResearchRepository struct {
	DB *gorm.DB
}

func NewResearchRepository(db *gorm.DB) *ResearchRepository {
	return &ResearchRepository{DB: db}
}

func (r *ResearchRepository) GetAllPublications() ([]models.Publication, error) {
	var publications []models.Publication
	err := r.DB.Find(&publications).Error
	return publications, err
}

func (r *ResearchRepository) CreatePublication(publication *models.Publication) error {
	return r.DB.Create(publication).Error
}

func (r *ResearchRepository) CreateConference(conference *models.Conference) error {
	return r.DB.Create(conference).Error
}

func (r *ResearchRepository) GetAllTheses() ([]models.Thesis, error) {
	var theses []models.Thesis
	err := r.DB.Preload("Supervisor").Find(&theses).Error
	return theses, err
}

func (r *ResearchRepository) CreateThesis(thesis *models.Thesis) error {
	return r.DB.Create(thesis).Error
}

func (r *ResearchRepository) AssignSupervisor(thesisID, supervisorID uint) error {
	var thesis models.Thesis
	if err := r.DB.First(&thesis, thesisID).Error; err != nil {
		return err
	}
	thesis.SupervisorID = supervisorID
	return r.DB.Save(&thesis).Error
}
