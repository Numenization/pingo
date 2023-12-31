package ui

import (
	"fmt"
	"image"
	"pingo/graph"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type PingoState struct {
	Graph *graph.PingoGraph

	// Main struct data
	running       bool
	interval      int
	pointsToGraph int
	target        string
	stopChan      chan bool
	canvasImage   image.Image
	logData       []string

	// Elements that are part of the UI
	window        fyne.Window
	canvasRaster  *canvas.Raster
	stopButton    *widget.Button
	startButton   *widget.Button
	saveButton    *widget.Button
	intervalEntry *NumericalEntry
	pointsEntry   *NumericalEntry
	targetEntry   *widget.Entry
	logGrid       *widget.TextGrid
	logScroll     *container.Scroll
}

func NewState() *PingoState {
	pGraph := &graph.PingoGraph{
		XValues: make([]float64, 0),
		YValues: make([]float64, 0),
		Length:  0,
	}

	state := &PingoState{
		running:       false,
		interval:      150,
		pointsToGraph: 25,
		target:        "8.8.8.8",
		Graph:         pGraph,
		stopChan:      make(chan bool),
	}

	return state
}

func (state *PingoState) SetWindow(win fyne.Window) {
	state.window = win
}

func (state *PingoState) SetImage(newImage image.Image) {
	state.canvasImage = newImage
	state.canvasRaster.Refresh()
}

func (state *PingoState) Log(str string) {
	state.logData = append(state.logData, str)
	state.logGrid.SetText(state.logGrid.Text() + str + "\n")
	state.logScroll.ScrollToBottom()
}

func (state *PingoState) String() string {
	return fmt.Sprintf(
		"state {%v}:\nrunning: %v\ninterval: %v\npoints to show: %v\ntarget: %v",
		&state,
		state.running,
		state.interval,
		state.pointsToGraph,
		state.target,
	)
}
