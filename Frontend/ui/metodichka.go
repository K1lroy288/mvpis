package ui

import (
	"fmt"
	"log"
	"net/url"
	"os/exec"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	"github.com/antonminin/borodyansky/api"
)

func BuildMetodichkaView(win fyne.Window, content *fyne.Container) fyne.CanvasObject {
	disciplinesBox := container.NewVBox()

	refreshButton := widget.NewButton("Обновить список дисциплин", func() {
		disciplines, err := api.GetAllDisciplines()
		if err != nil {
			disciplinesBox.Add(widget.NewLabel("Ошибка загрузки дисциплин"))
			return
		}
		disciplinesBox.Objects = nil // Очищаем предыдущие элементы

		for _, discipline := range disciplines {
			// Кнопка для дисциплины
			disciplineButton := widget.NewButton(discipline.Name, func() {
				// Создаём контейнер для дополнительных опций (загрузка файла)
				optionsContainer := container.NewVBox()

				// Кнопка загрузки файла
				uploadButton := widget.NewButton("Загрузить файл", func() {
					// Диалог выбора файла
					fd := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
						if err != nil {
							dialog.ShowError(err, win)
							return
						}
						if closer == nil {
							return
						}
						//Здесь будет логика загрузки
						fileURL := closer.URI().String()
						content.Objects = []fyne.CanvasObject{widget.NewLabel(fmt.Sprintf("Загрузка файла: %s для дисциплины: %s", fileURL, discipline.Name))}
						content.Refresh()
						log.Println(fileURL)
						// TODO: загрузка файла
						// API call
						uploadFileToDiscipline(win, discipline.ID, closer.URI())
					}, win)

					fd.SetFilter(storage.NewExtensionFileFilter([]string{".pdf", ".txt", ".docx"})) // Фильтр файлов
					fd.Show()
				})
				optionsContainer.Add(uploadButton)

				// Кнопка получения файлов
				getFilesButton := widget.NewButton("Получить файлы", func() {
					// API call to get files for this discipline
					files, err := api.GetFilesForDiscipline(discipline.ID)
					if err != nil {
						dialog.ShowError(err, win)
						return
					}

					// Display the files
					filesContainer := container.NewVBox()
					if len(files) == 0 {
						filesContainer.Add(widget.NewLabel("Нет файлов для этой дисциплины."))
					} else {
						for _, file := range files {
							fileButton := widget.NewButton(file.Name, func() {
								// Open the file using the default application
								err := openFile(file.Path)
								if err != nil {
									dialog.ShowError(fmt.Errorf("Не удалось открыть файл: %s", err), win)
								}
							})
							filesContainer.Add(fileButton)
						}
					}

					content.Objects = []fyne.CanvasObject{filesContainer}
					content.Refresh()
				})
				optionsContainer.Add(getFilesButton)

				// Отображаем контейнер с опциями
				content.Objects = []fyne.CanvasObject{optionsContainer}
				content.Refresh()
			})
			disciplinesBox.Add(disciplineButton)
		}
		disciplinesBox.Refresh()
	})

	return container.NewVBox(
		widget.NewLabel("Учебно-методическая деятельность"),
		refreshButton,
		disciplinesBox,
	)
}

func uploadFileToDiscipline(win fyne.Window, disciplineID uint, fileURI fyne.URI) {
	// Реализация загрузки файла на сервер
	// Открыть файл
	r, err := storage.OpenFileFromURI(fileURI)
	if err != nil {
		log.Println(err)
		dialog.ShowError(err, win)
		return
	}
	defer r.Close()
	// POST запрос

	// Create a URL from the string
	uploadURL := fmt.Sprintf("http://localhost:8081/api/v1/methodology/disciplines/%d/files", disciplineID)
	u, err := url.Parse(uploadURL)
	if err != nil {
		dialog.ShowError(err, win)
		return
	}
	err = api.SendFile(u, r, fileURI.Name(), win)
	if err != nil {
		dialog.ShowError(err, win)
		return
	}
}

func openFile(filePath string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", filePath)
	case "darwin":
		cmd = exec.Command("open", filePath)
	case "linux":
		cmd = exec.Command("kate", filePath)
	default:
		return fmt.Errorf("unsupported operating system")
	}
	return cmd.Start()
}
