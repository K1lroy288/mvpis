package repositories

import (
	"educational-service/models"

	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

func (r *StudentRepository) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	err := r.DB.Preload("Grades").Find(&students).Error
	return students, err
}

func (r *StudentRepository) GetStudentByID(id uint) (models.Student, error) {
	var student models.Student
	err := r.DB.Preload("Grades").Where("id = ?", id).First(&student).Error
	return student, err
}

func (r *StudentRepository) CreateStudent(student *models.Student) error {
	return r.DB.Create(student).Error
}

func (r *StudentRepository) AddGrade(grade *models.Grade) error {
	return r.DB.Create(grade).Error
}

func (r *StudentRepository) GetHonorStudents() ([]models.Student, error) {
	var students []models.Student
	err := r.DB.Raw(`
		SELECT s.*, COALESCE(AVG(g.value), 0) as avg_grade
		FROM students s
		LEFT JOIN grades g ON s.id = g.student_id
		GROUP BY s.id
		HAVING COALESCE(AVG(g.value), 0) >= ?
	`, 4.5).Preload("Grades").Scan(&students).Error

	return students, err
}

func (r *StudentRepository) GetExpelledStudents() ([]models.Student, error) {
	var students []models.Student
	err := r.DB.Raw(`
		SELECT s.*, COALESCE(AVG(g.value), 0) as avg_grade
		FROM students s
		LEFT JOIN grades g ON s.id = g.student_id
		GROUP BY s.id
		HAVING COALESCE(AVG(g.value), 0) <= ?
	`, 2.5).Preload("Grades").Scan(&students).Error

	return students, err
}
