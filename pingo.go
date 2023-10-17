package main

import (
	"pingo/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	newApp := app.New()
	state := ui.NewState()

	window := ui.CreateWindow(newApp, state)
	window.Resize(fyne.NewSize(900, 190))

	window.ShowAndRun()
}
