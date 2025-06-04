package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type Student struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Group    string  `json:"group"`
	AvgGrade float64 `json:"avg_grade"`
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

func SendFile(uploadURL *url.URL, fileReader io.Reader, fileName string, win fyne.Window) error {
	// Create a new multipart form
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Create a form file field
	part, err := writer.CreateFormFile("file", filepath.Base(fileName))
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Copy the file data to the form file
	_, err = io.Copy(part, fileReader)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create a new request
	req, err := http.NewRequest("POST", uploadURL.String(), body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("ошибка: %s", string(body))
	}
	// Read the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(responseBody))
	dialog.ShowInformation("Успех", "Файл успешно загружен", win)
	return nil
}

type File struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	DisciplineID uint   `json:"discipline_id"`
}

func GetFilesForDiscipline(disciplineID uint) ([]File, error) {
	// Create the URL
	url := fmt.Sprintf("http://localhost:8081/api/v1/methodology/disciplines/%d/files", disciplineID)

	// Send the GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response
	var files []File
	err = json.Unmarshal(body, &files)
	if err != nil {
		return nil, err
	}

	return files, nil
}
