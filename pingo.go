package main

import (
	"fmt"
	"math/rand"
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
		x := 20

		state.Graph.Length = x
		for i := 0; i < x; i++ {
			state.Graph.AddValue(float64(rand.Intn(100)))
		}
		img, err := state.Graph.GenerateImage()
		if err != nil {
			fmt.Println(err)
			return
		}
		state.SetImage(img)

		time.Sleep(5 * time.Second)

		fmt.Println("Trying to change image!")

		state.Graph.Clear()
		for i := 0; i < x; i++ {
			state.Graph.AddValue(float64(rand.Intn(100)))
		}
		img2, err2 := state.Graph.GenerateImage()
		if err2 != nil {
			fmt.Println(err)
			return
		}
		state.SetImage(img2)

	}()

	window.ShowAndRun()
}
