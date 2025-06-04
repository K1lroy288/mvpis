package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildUchebna() *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Учёт успеваемости студентов"),
		widget.NewButton("Загрузить оценки (API)", func() {
			// Тут делаем запрос к API
		}),
		widget.NewButton("Показать аналитику", func() {
			// Подсчёт ср.балла и т.п.
		}),
	)
}
