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
	conferencesBox := container.NewVBox()

	// Function to refresh publications
	refreshPublications := func() {
		publicationsBox.Objects = nil // Clear existing publications
		publications, err := api.GetAllPublications()
		if err != nil {
			publicationsBox.Add(widget.NewLabel("Ошибка загрузки публикаций"))
			publicationsBox.Refresh()
			return
		}
		for _, publication := range publications {
			label := widget.NewLabel(fmt.Sprintf("%s - %s, %s (%d)", publication.Title, publication.Authors, publication.Journal, publication.Year))
			publicationsBox.Add(label)
		}
		publicationsBox.Refresh()
	}

	// Function to refresh conferences
	refreshConferences := func() {
		conferencesBox.Objects = nil // Clear existing conferences
		conferences, err := api.GetAllConferences()
		if err != nil {
			conferencesBox.Add(widget.NewLabel("Ошибка загрузки конференций"))
			conferencesBox.Refresh()
			return
		}
		for _, conference := range conferences {
			label := widget.NewLabel(fmt.Sprintf("%s - %s", conference.Name, conference.Date))
			conferencesBox.Add(label)
		}
		conferencesBox.Refresh()
	}

	// Conference registration form
	nameEntry := widget.NewEntry()
	dateEntry := widget.NewEntry()
	participantEntry := widget.NewEntry()

	registrationForm := container.NewVBox(
		widget.NewLabel("Регистрация конференции"),
		widget.NewLabel("Название:"), nameEntry,
		widget.NewLabel("Дата:"), dateEntry,
		widget.NewLabel("Участник:"), participantEntry,
	)
	registrationForm.Hide() // Initially hide the form

	registerConferenceBtn := widget.NewButton("Зарегистрировать конференцию", func() {
		// Toggle the visibility of the registration form
		if registrationForm.Visible() {
			registrationForm.Hide()
		} else {
			registrationForm.Show()
		}

		conference := api.Conference{
			Name:        nameEntry.Text,
			Date:        dateEntry.Text,
			Participant: participantEntry.Text,
		}
		err := api.CreateConference(conference)
		if err != nil {
			widget.ShowPopUp(widget.NewLabel(fmt.Sprintf("Ошибка регистрации: %s", err.Error())), fyne.CurrentApp().Driver().CanvasForObject(content))
			return
		}
		widget.ShowPopUp(widget.NewLabel("Конференция успешно зарегистрирована"), fyne.CurrentApp().Driver().CanvasForObject(content))
		// Optionally, refresh the conference list after registration
		// refreshConferences()
	})

	// Buttons
	getPublicationsBtn := widget.NewButton("Получить список публикаций", refreshPublications)
	getConferencesBtn := widget.NewButton("Получить список конференций", refreshConferences)

	return container.NewVBox(
		widget.NewLabel("Научная деятельность"),
		getPublicationsBtn,
		publicationsBox,
		widget.NewLabel("Конференции"),
		getConferencesBtn,
		conferencesBox,
		registerConferenceBtn, // Add the button to toggle the form
		registrationForm,      // Add the registration form to the layout
	)
}
