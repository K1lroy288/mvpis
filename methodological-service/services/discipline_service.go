package services

import (
	"methodological-service/models"
	"methodological-service/repositories"
)

type DisciplineService struct {
	repo *repositories.DisciplineRepository
}

func NewDisciplineService(repo *repositories.DisciplineRepository) *DisciplineService {
	return &DisciplineService{repo: repo}
}

func (s *DisciplineService) GetAllDisciplines() ([]models.Discipline, error) {
	return s.repo.GetAllDisciplines()
}

func (s *DisciplineService) GetDisciplineByID(id uint) (models.Discipline, error) {
	return s.repo.GetDisciplineByID(id)
}

func (s *DisciplineService) RegisterDiscipline(discipline *models.Discipline) error {
	return s.repo.CreateDiscipline(discipline)
}

func (s *DisciplineService) UploadFile(file *models.File) error {
	return s.repo.AddFile(file)
}
