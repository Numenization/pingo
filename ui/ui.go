package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func UpdateGraph() {

}

func CreateWindow(a fyne.App) fyne.Window {
	window := a.NewWindow("Pingo")

	// Large row with just the output graph
	//chartRow := container.New(layout.NewHBoxLayout())
	img := canvas.NewImageFromFile("resouce/fynetest.png")
	img.FillMode = canvas.ImageFillOriginal

	controls := BuildControlsLayout()

	bottomContainer := container.NewGridWithColumns(2, controls, widget.NewEntry())

	mainContainer := container.NewBorder(img, bottomContainer, nil, nil)

	window.SetContent(mainContainer)

	return window
}
