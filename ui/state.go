package ui

import (
	"image"
	"pingo/graph"

	"fyne.io/fyne/v2/canvas"
)

type PingoState struct {
	running       bool
	interval      int
	pointsToGraph int
	Graph         *graph.PingoGraph
	canvasImage   image.Image
	canvasRaster  *canvas.Raster
	stopChan      chan bool
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
		Graph:         pGraph,
		stopChan:      make(chan bool),
	}

	return state
}

func (state *PingoState) SetImage(newImage image.Image) {
	state.canvasImage = newImage
	state.canvasRaster.Refresh()
}
