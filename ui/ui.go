package ui

import (
	"fmt"
	"time"

	probing "github.com/prometheus-community/pro-bing"
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

	if state.pointsToGraph <= 2 || state.pointsToGraph > 1000 {
		return &StateError{
			str: fmt.Sprintf("ui: invalid max graph points %v", state.pointsToGraph),
		}
	}

	state.Graph.Clear()
	state.Graph.Length = state.pointsToGraph

	pinger, err := probing.NewPinger("www.google.com")

	if err != nil {
		return err
	}

	pinger.Interval = time.Duration(state.interval) * time.Millisecond

	pinger.OnRecv = func(pkt *probing.Packet) {
		state.Graph.AddValue(float64(pkt.Rtt / time.Millisecond))
	}

	go func() {
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
				//state.Graph.AddValue(float64(rand.Intn(100)))
				img, err := state.Graph.GenerateImage()
				if err != nil {
					//fmt.Println(err)
					//running = false
					continue
				}
				state.SetImage(img)
				time.Sleep(time.Duration(state.interval) * time.Millisecond)
			}
		}
	}()

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
	}
}
