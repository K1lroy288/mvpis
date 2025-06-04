package services

import (
	"educational-service/models"
	"educational-service/repositories"
)

type StudentService struct {
	repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) GetAllStudents() ([]models.Student, error) {
	students, err := s.repo.GetAllStudents()
	if err != nil {
		return nil, err
	}

	// Рассчитываем средний балл для каждого студента
	for i, student := range students {
		var sum float64
		for _, grade := range student.Grades {
			sum += grade.Value
		}
		if len(student.Grades) > 0 {
			students[i].AvgGrade = sum / float64(len(student.Grades))
		} else {
			students[i].AvgGrade = 0
		}
	}

	return students, nil
}

func (s *StudentService) GetStudentByID(id uint) (models.Student, error) {
	return s.repo.GetStudentByID(id)
}

func (s *StudentService) RegisterStudent(student *models.Student) error {
	return s.repo.CreateStudent(student)
}

func (s *StudentService) RecordGrade(grade *models.Grade) error {
	return s.repo.AddGrade(grade)
}
