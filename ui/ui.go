package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/dialog"
	probing "github.com/prometheus-community/pro-bing"
)

// Starts graphing the Pingo Graph
func StartGraphLoop(state *PingoState) error {
	if state.running {
		return &StateError{
			str: "Graphing loop already running",
		}
	}

	if state.interval < 50 {
		return &StateError{
			str: fmt.Sprintf("Invalid update interval %vms (must be greater than 50ms)", state.interval),
		}
	}

	if state.pointsToGraph <= 2 || state.pointsToGraph > 1000 {
		return &StateError{
			str: fmt.Sprintf("Invalid max graph points %v (must be between 2 and 999)", state.pointsToGraph),
		}
	}

	targetUrl := state.target

	state.Graph.Clear()
	state.Graph.Length = state.pointsToGraph

	pinger, err := probing.NewPinger(targetUrl)

	if err != nil {
		return err
	}

	pinger.Interval = time.Duration(state.interval) * time.Millisecond

	pinger.OnRecv = func(pkt *probing.Packet) {
		state.Graph.AddValue(float64(pkt.Rtt / time.Millisecond))
	}

	go GraphLoop(state, pinger)
	state.Log(fmt.Sprintf("Started pinging '%v' with an interval of %vms and graphing %v points", state.target, state.interval, state.pointsToGraph))

	err = pinger.Run()

	if err != nil {
		return err
	}

	return nil
}

// Stops the Pingo graphing loop
func StopGraphLoop(state *PingoState) {
	if state.running {
		state.stopChan <- true
		state.Log("Stopped pinging")
	}
}

func GraphLoop(state *PingoState, pinger *probing.Pinger) {
	state.running = true
	for {
		if !state.running {
			break
		}
		select {
		case <-state.stopChan:
			pinger.Stop()
			state.running = false
		default:
			img, err := state.Graph.GenerateImage(state.canvasRaster.Size())
			if err != nil {
				state.running = false
				pinger.Stop()
				dialog.ShowError(err, state.window)
				continue
			}
			state.SetImage(img)
			time.Sleep(time.Duration(state.interval) * time.Millisecond)
		}
	}
}
