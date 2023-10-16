package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BuildControlsLayout() *fyne.Container {
	// Box containing all the user controlled fields
	// ------------------------------ |
	// Points to Show:       [    25] |
	// Ping Interval (ms):   [   150] |
	// Target IPv4:   [      8.8.8.8] |     Logs go here, fills horizontal space
	// (    Start    ) (    Stop    ) |
	// (         Save  Logs         ) |
	// ------------------------------ |
	pointsLabel := widget.NewLabel("Points to Show:")
	intervalLabel := widget.NewLabel("Ping Interval (ms):")
	targetLabel := widget.NewLabel("Target IPv4:")

	pointsEntry := widget.NewEntry()
	pointsEntry.SetText("25")
	intervalEntry := widget.NewEntry()
	intervalEntry.SetText("150")
	targetEntry := widget.NewEntry()
	targetEntry.SetText("8.8.8.8")

	startButton := widget.NewButton("Start", func() {})
	stopButton := widget.NewButton("Stop", func() {})
	saveButton := widget.NewButton("Save Logs", func() {})

	controlForm := container.New(layout.NewFormLayout(), pointsLabel, pointsEntry, intervalLabel, intervalEntry, targetLabel, targetEntry)
	startStopButtons := container.New(layout.NewGridLayout(2), startButton, stopButton)

	spacer := layout.NewSpacer()
	spacer.Resize(fyne.NewSize(300, 0))
	fmt.Println(spacer.MinSize())

	controls := container.New(layout.NewVBoxLayout(), controlForm, startStopButtons, saveButton, spacer)
	return controls
}

func CreateWindow(a fyne.App) fyne.Window {
	window := a.NewWindow("Pingo")

	// Large row with just the output graph
	//chartRow := container.New(layout.NewHBoxLayout())

	controls := BuildControlsLayout()

	bottomContainer := container.NewGridWithColumns(2, controls, widget.NewEntry())

	window.SetContent(bottomContainer)

	return window
}
