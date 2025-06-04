package models

import "gorm.io/gorm"

type Publication struct {
	gorm.Model
	Title    string `json:"title"`
	Authors  string `json:"authors"`
	Journal  string `json:"journal"`
	Year     int    `json:"year"`
}

type Conference struct {
	gorm.Model
	Name      string `json:"name"`
	Date      string `json:"date"`
	Participant string `json:"participant"`
}

type Thesis struct {
	gorm.Model
	Topic       string `json:"topic"`
	StudentName string `json:"student_name"`
	Year        int    `json:"year"`
	SupervisorID uint   `json:"supervisor_id"`
	Supervisor   Supervisor `gorm:"foreignKey:SupervisorID"`
}

type Supervisor struct {
	gorm.Model
	Name string `json:"name"`
}