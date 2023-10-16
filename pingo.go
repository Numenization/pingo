package main

import (
	"pingo/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := ui.CreateWindow(a)
	w.Resize(fyne.NewSize(900, 190))

	w.ShowAndRun()
}
