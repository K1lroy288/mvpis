package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

func GetHonorStudents() ([]Student, error) {
	resp, err := http.Get("http://localhost:8080/api/v1/students/honor")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var students []Student
	err = json.Unmarshal(body, &students)
	log.Printf("student: %+v", students[0])
	return students, err
}

func GetExpelledStudents() ([]Student, error) {
	resp, err := http.Get("http://localhost:8080/api/v1/students/expelled")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var students []Student
	err = json.Unmarshal(body, &students)
	log.Printf("student: %+v", students[0])
	return students, err
}
