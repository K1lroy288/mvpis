package ui

import (
	"fmt"

	"github.com/antonminin/borodyansky/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildEducationView() fyne.CanvasObject {
	studentsBox := container.NewVBox()

	// Pagination variables
	pageSize := 10
	currentPage := 0
	var allStudents []api.Student // Store all students
	var currentStudents []api.Student

	// Function to display students with pagination
	displayStudents := func(students []api.Student, title string) {
		studentsBox.Objects = nil // Clear existing content
		studentsBox.Add(widget.NewLabel(title))
		if len(students) == 0 {
			studentsBox.Add(widget.NewLabel("Нет данных"))
		}
		// Calculate start and end indices for pagination
		start := currentPage * pageSize
		end := start + pageSize
		if end > len(students) {
			end = len(students)
		}
		currentStudents = students[start:end] // Get students for the current page

		for _, s := range currentStudents {
			fmt.Println(s.AvgGrade) // Print the value of AvgGrade
			studentsBox.Add(widget.NewLabel(s.Name + " (" + s.Group + ") - Средний балл: " + fmt.Sprintf("%.2f", s.AvgGrade)))
		}
		studentsBox.Refresh()
	}

	// Buttons
	honorBtn := widget.NewButton("Отличники", func() {
		currentPage = 0 // Reset to first page
		students, err := api.GetHonorStudents()
		if err != nil {
			studentsBox.Objects = []fyne.CanvasObject{widget.NewLabel("Ошибка загрузки отличников")}
			studentsBox.Refresh()
			return
		}
		allStudents = students // Store the students
		displayStudents(allStudents, "Список отличников:")

	})

	expelledBtn := widget.NewButton("На отчисление", func() {
		currentPage = 0 // Reset to first page
		students, err := api.GetExpelledStudents()
		if err != nil {
			studentsBox.Objects = []fyne.CanvasObject{widget.NewLabel("Ошибка загрузки студентов на отчисление")}
			studentsBox.Refresh()
			return
		}
		allStudents = students // Store the students
		displayStudents(allStudents, "Список на отчисление:")
	})

	//pagination := container.NewHBox() // Remove pagination buttons

	content := container.NewBorder(
		container.NewVBox(honorBtn, expelledBtn), // Top: Buttons
		nil,                                      // Bottom: Pagination
		nil,
		nil,
		studentsBox, // Center: Student List
	)

	scrollContainer := container.NewScroll(content)
	return scrollContainer
}
