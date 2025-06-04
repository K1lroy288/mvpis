package ui

import (
	"fmt"

	"github.com/antonminin/borodyansky/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildMetodichkaView(content *fyne.Container) fyne.CanvasObject {
	disciplinesBox := container.NewVBox()

	disciplines, err := api.GetAllDisciplines()
	if err != nil {
		disciplinesBox.Add(widget.NewLabel("Ошибка загрузки дисциплин"))
		return disciplinesBox
	}

	for _, discipline := range disciplines {
		btn := widget.NewButton(discipline.Name, func() {
			// Handle button click (e.g., show discipline details)
			content.Objects = []fyne.CanvasObject{widget.NewLabel(fmt.Sprintf("Выбрана дисциплина: %s", discipline.Name))}
			content.Refresh()
		})
		disciplinesBox.Add(btn)
	}

	return container.NewVBox(
		widget.NewLabel("Учебно-методическая деятельность"),
		disciplinesBox,
	)
}
