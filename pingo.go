package main

import (
	"fmt"
	"pingo/ui"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	newApp := app.New()
	state := ui.NewState()

	window := ui.CreateWindow(newApp, state)
	window.Resize(fyne.NewSize(900, 190))

	go func() {
		time.Sleep(5 * time.Second)

		fmt.Println("Trying to change image!")

	}()

	window.ShowAndRun()
}
