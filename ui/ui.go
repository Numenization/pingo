package ui

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Starts graphing the Pingo Graph
func StartGraphLoop(state *PingoState) error {
	if state.running {
		return &StateError{
			str: "ui: graphing loop already running",
		}
	}

	if state.interval < 50 {
		return &StateError{
			str: fmt.Sprintf("ui: invalid update interval %v", state.interval),
		}
	}

	go func() {
		running := true
		for {
			if !running {
				break
			}
			select {
			case <-state.stopChan:
				running = false
			default:
				state.Graph.GenerateImage()
				time.Sleep(time.Duration(state.interval) * time.Millisecond)
			}
		}
	}()
	return nil
}

func CreateWindow(a fyne.App, state *PingoState) fyne.Window {
	window := a.NewWindow("Pingo")

	// Large row with just the output graph
	//chartRow := container.New(layout.NewHBoxLayout())

	raster := canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		if state.canvasImage != nil {
			return state.canvasImage.At(x, y)
		} else {
			return color.RGBA{0, 0, 0, 0}
		}
	})
	state.canvasRaster = raster
	raster.SetMinSize(fyne.NewSize(900, 190))
	//img.FillMode = canvas.ImageFillOriginal

	controls := BuildControlsLayout()

	bottomContainer := container.NewGridWithColumns(2, controls, widget.NewEntry())

	mainContainer := container.NewBorder(raster, bottomContainer, nil, nil)

	window.SetContent(mainContainer)

	return window
}
