package models

type Student struct {
	Name     string  `json:"name"`
	Group    string  `json:"group"`
	Grades   []Grade `json:"grades" `
	AvgGrade float64 `json:"avg_grade"` // поле не сохраняется в БД
}

type Grade struct {
	Subject   string  `json:"subject"`
	Value     float64 `json:"value"`
	Semester  int     `json:"semester"`
	StudentID uint
}
