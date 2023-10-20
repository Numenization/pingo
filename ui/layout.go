package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Layout that will create an evenly sized horizontal grid and can be configured to allow a child object to span across 2 or more grid spaces
type HorizSpanLayout struct {
	spans            []int
	biggestMinHeight float32
}

func NewHorizSpanLayout(s []int) *HorizSpanLayout {
	return &HorizSpanLayout{
		spans: s,
	}
}

func (l *HorizSpanLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	gridTotal := 0
	for _, num := range l.spans {
		gridTotal += num
	}

	// dont divide by 0
	if len(l.spans) <= 0 || gridTotal <= 0 {
		return
	}

	gridSize := containerSize.Width / float32(gridTotal)
	pos := fyne.NewPos(0, 0)

	for i, o := range objects {
		span := l.spans[i]
		w, h := gridSize*float32(span), o.MinSize().Height

		if h > l.biggestMinHeight {
			l.biggestMinHeight = h
		} else {
			h = l.biggestMinHeight
		}

		size := fyne.NewSize(w, h)

		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width, 0))
	}
}

func (l *HorizSpanLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		size := o.MinSize()

		w += size.Width
		h += size.Height
	}
	return fyne.NewSize(w, h)
}

func BuildRaster(state *PingoState) *fyne.Container {
	raster := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		if state.canvasImage != nil {
			return state.canvasImage.At(x, y)
		} else {
			return color.RGBA{255, 255, 255, 255}
		}
	})
	state.canvasRaster = raster
	raster.SetMinSize(fyne.NewSize(1000, 190))

	rasterContainer := container.NewStack(raster)
	return rasterContainer
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

	startButton := widget.NewButton("Start", func() {
		go func() {
			err := StartGraphLoop(state)
			if err != nil {
				dialog.ShowError(err, state.window)
			}
		}()
	})
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
	state.SetWindow(window)

	raster := BuildRaster(state)
	controls := BuildControlsLayout(state)

	logGrid := widget.NewTextGrid()
	logScroll := container.NewScroll(logGrid)

	logGrid.ShowLineNumbers = true

	state.logGrid = logGrid
	state.logScroll = logScroll

	bottomLayout := NewHorizSpanLayout([]int{1, 3})
	bottomContainer := container.New(bottomLayout, controls, logScroll)
	mainContainer := container.NewBorder(nil, bottomContainer, nil, nil, raster)

	window.SetContent(mainContainer)
	return window
}
