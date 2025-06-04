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
	refreshBtn := widget.NewButton("Загрузить студентов", func() {
		studentsBox.Objects = nil
		students, err := api.GetAllStudents()
		if err != nil {
			studentsBox.Add(widget.NewLabel("Ошибка загрузки"))
			return
		}
		for _, s := range students {
			studentsBox.Add(widget.NewLabel(s.Name + " (" + s.Group + ") - Средний балл: " + fmt.Sprintf("%.2f", s.AvgGrade))) // Modified line
		}
		studentsBox.Refresh()
	})

	return container.NewVBox(
		widget.NewLabel("Учебная деятельность"),
		refreshBtn,
		studentsBox,
	)
}
