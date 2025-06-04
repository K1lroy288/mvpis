package repositories

import (
	"methodological-service/models"

	"gorm.io/gorm"
)

type DisciplineRepository struct {
	DB *gorm.DB
}

func NewDisciplineRepository(db *gorm.DB) *DisciplineRepository {
	return &DisciplineRepository{DB: db}
}

func (r *DisciplineRepository) GetAllDisciplines() ([]models.Discipline, error) {
	var disciplines []models.Discipline
	err := r.DB.Preload("Files").Find(&disciplines).Error
	return disciplines, err
}

func (r *DisciplineRepository) GetDisciplineByID(id uint) (models.Discipline, error) {
	var discipline models.Discipline
	err := r.DB.Preload("Files").Where("id = ?", id).First(&discipline).Error
	return discipline, err
}

func (r *DisciplineRepository) CreateDiscipline(discipline *models.Discipline) error {
	return r.DB.Create(discipline).Error
}

func (r *DisciplineRepository) AddFile(file *models.File) error {
	return r.DB.Create(file).Error
}
