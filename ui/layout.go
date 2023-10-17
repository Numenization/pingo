package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BuildRaster(state *PingoState) *canvas.Raster {
	raster := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		if state.canvasImage != nil {
			return state.canvasImage.At(x, y)
		} else {
			return color.RGBA{0, 0, 0, 0}
		}
	})
	state.canvasRaster = raster
	raster.SetMinSize(fyne.NewSize(900, 190))

	return raster
}

func BuildControlsLayout(state *PingoState) *fyne.Container {
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

	boundPoints := binding.IntToString(binding.BindInt(&state.pointsToGraph))
	boundInterval := binding.IntToString(binding.BindInt(&state.interval))
	boundTarget := binding.BindString(&state.target)

	pointsEntry := NewNumericalEntryWithData(boundPoints)
	intervalEntry := NewNumericalEntryWithData(boundInterval)
	targetEntry := widget.NewEntryWithData(boundTarget)

	startButton := widget.NewButton("Start", func() { go func() { StartGraphLoop(state) }() })
	stopButton := widget.NewButton("Stop", func() { go func() { StopGraphLoop(state) }() })
	saveButton := widget.NewButton("Save Logs", func() { fmt.Println(state) })

	state.pointsEntry = pointsEntry
	state.intervalEntry = intervalEntry
	state.targetEntry = targetEntry
	state.startButton = startButton
	state.stopButton = stopButton
	state.saveButton = saveButton

	controlForm := container.New(layout.NewFormLayout(), pointsLabel, pointsEntry, intervalLabel, intervalEntry, targetLabel, targetEntry)
	startStopButtons := container.New(layout.NewGridLayout(2), startButton, stopButton)

	spacer := layout.NewSpacer()
	spacer.Resize(fyne.NewSize(300, 0))

	controls := container.New(layout.NewVBoxLayout(), controlForm, startStopButtons, saveButton, spacer)
	return controls
}

func CreateWindow(a fyne.App, state *PingoState) fyne.Window {
	window := a.NewWindow("Pingo")

	raster := BuildRaster(state)
	controls := BuildControlsLayout(state)

	bottomContainer := container.NewGridWithColumns(2, controls, widget.NewEntry())
	mainContainer := container.NewBorder(raster, bottomContainer, nil, nil)

	window.SetContent(mainContainer)

	return window
}
