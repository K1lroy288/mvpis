package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/antonminin/borodyansky/api"
)

func BuildNaukaView(content *fyne.Container) fyne.CanvasObject {
	publicationsBox := container.NewVBox()

	publications, err := api.GetAllPublications()
	if err != nil {
		publicationsBox.Add(widget.NewLabel("Ошибка загрузки публикаций"))
		return publicationsBox
	}

	for _, publication := range publications {
		label := widget.NewLabel(fmt.Sprintf("%s - %s, %s (%d)", publication.Title, publication.Authors, publication.Journal, publication.Year))
		publicationsBox.Add(label)
	}

	return container.NewVBox(
		widget.NewLabel("Научная деятельность"),
		publicationsBox,
	)
}
