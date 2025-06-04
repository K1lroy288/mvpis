// ui/gia.go
package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildGIAView() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Государственная Итоговая Аттестация (ГИА)"),
		widget.NewButton("Регистрация темы дипломной работы", func() {
			widget.NewLabel("Тема зарегистрирована (имитация)")
		}),
		widget.NewButton("Привязка к научному руководителю", func() {
			widget.NewLabel("Научный руководитель назначен (имитация)")
		}),
	)
}
