package services

import (
	"educational-service/models"
	"educational-service/repositories"
)

type ResearchService struct {
	repo *repositories.ResearchRepository
}

func NewResearchService(repo *repositories.ResearchRepository) *ResearchService {
	return &ResearchService{repo: repo}
}

func (s *ResearchService) GetAllPublications() ([]models.Publication, error) {
	return s.repo.GetAllPublications()
}

func (s *ResearchService) CreatePublication(publication *models.Publication) error {
	return s.repo.CreatePublication(publication)
}

func (s *ResearchService) CreateConference(conference *models.Conference) error {
	return s.repo.CreateConference(conference)
}

func (s *ResearchService) GetAllTheses() ([]models.Thesis, error) {
	return s.repo.GetAllTheses()
}

func (s *ResearchService) CreateThesis(thesis *models.Thesis) error {
	return s.repo.CreateThesis(thesis)
}

func (s *ResearchService) AssignSupervisor(thesisID, supervisorID uint) error {
	return s.repo.AssignSupervisor(thesisID, supervisorID)
}