package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/antonminin/borodyansky/ui" // Исправим этот импорт на следующем шаге
)

func main() {
	myApp := app.New()                                    // Создаём новое GUI-приложение
	myWindow := myApp.NewWindow("Университетский клиент") // Создаём главное окно
	myWindow.Resize(fyne.NewSize(1000, 600))

	content := container.NewMax()
	sidebar := ui.BuildSidebar(content)

	mainContainer := container.NewHSplit(sidebar, content) // Горизонтальный сплит: слева сайдбар, справа контент
	mainContainer.Offset = 0.2                             // Ширина сайдбара (20%)

	myWindow.SetContent(mainContainer) // Устанавливаем содержимое окна
	myWindow.ShowAndRun()              // Показываем и запускаем приложение
}
