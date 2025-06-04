package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BuildSidebar(content *fyne.Container) *fyne.Container {
	return container.NewVBox(
		widget.NewLabel("Модули"),
		widget.NewButton("Учебная деятельность", func() {
			content.Objects = []fyne.CanvasObject{BuildEducationView()}
			content.Refresh()
		}),
		widget.NewButton("Учебно-методическая", func() {
			content.Objects = []fyne.CanvasObject{BuildMetodichkaView(content)}
			content.Refresh()
		}),
		widget.NewButton("Научная деятельность", func() {
			content.Objects = []fyne.CanvasObject{BuildNaukaView(content)}
			content.Refresh()
		}),
		layout.NewSpacer(),
	)
}
