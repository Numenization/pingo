package graph

import (
	"fmt"
	"image"

	"github.com/wcharczuk/go-chart/v2"
)

type PingoGraph struct {
	XValues []float64
	YValues []float64

	Length int
}

// Sets the maximum length of a Pingo Graph. Valid lengths are all integers 0 < x < 1000
func (graph *PingoGraph) SetLength(length int) bool {
	if length < 0 && length > 1000 {
		return false
	}

	graph.Length = length
	return true
}

// Add a value to the graph. If the graph is already at max length, remove the first value in the graph
func (graph *PingoGraph) AddValue(val float64) {
	if len(graph.YValues) >= graph.Length {
		// Add the new value to the end and get rid of the first value
		graph.YValues = append(graph.YValues, val)[1:]
	} else {
		// Append new value to YValues and add a new index to XValues
		graph.YValues = append(graph.YValues, val)
		graph.XValues = append(graph.XValues, float64(len(graph.XValues)))
	}
}

// Clear all data from the graph
func (graph *PingoGraph) Clear() {
	graph.YValues = make([]float64, 0)
	graph.XValues = make([]float64, 0)
}

// Creates an Image based on the current data of the graph
func (graph *PingoGraph) GenerateImage() image.Image {
	gochart := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: graph.XValues,
				YValues: graph.YValues,
			},
		},
	}

	collector := &chart.ImageWriter{}
	gochart.Render(chart.PNG, collector)

	img, err := collector.Image()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return img
}
