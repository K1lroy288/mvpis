package models

import "gorm.io/gorm"

type Discipline struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
	Files       []File `json:"files" gorm:"foreignKey:DisciplineID"`
}

type File struct {
	gorm.Model
	Name         string `json:"name"`
	Path         string `json:"path"` // путь к файлу или URL
	DisciplineID uint
}
