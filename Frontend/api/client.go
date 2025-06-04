package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Student struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Group    string  `json:"group"`
	AvgGrade float64 `json:"avg_grade"` // Add this line
}

type Grade struct {
	StudentID uint    `json:"student_id"`
	Subject   string  `json:"subject"`
	Score     float64 `json:"score"`
	Passed    bool    `json:"passed"`
}

func GetAllStudents() ([]Student, error) {
	resp, err := http.Get("http://localhost:8080/api/v1/students/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var students []Student
	err = json.Unmarshal(body, &students)
	return students, err
}

func AddGrade(grade Grade) error {
	data, _ := json.Marshal(grade)
	resp, err := http.Post("http://localhost:8080/api/v1/students/"+fmt.Sprint(grade.StudentID)+"/grades",
		"application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("ошибка: %s", string(body))
	}
	return nil
}

type Discipline struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetAllDisciplines() ([]Discipline, error) {
	resp, err := http.Get("http://localhost:8081/api/v1/methodology/disciplines")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var disciplines []Discipline
	err = json.Unmarshal(body, &disciplines)
	return disciplines, err
}

type Publication struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Journal string `json:"journal"`
	Year    int    `json:"year"`
}

func GetAllPublications() ([]Publication, error) {
	resp, err := http.Get("http://localhost:8082/api/v1/research/publications")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var publications []Publication
	err = json.Unmarshal(body, &publications)
	return publications, err
}
