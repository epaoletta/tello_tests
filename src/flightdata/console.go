package flightdata

import (
	"log"

	"github.com/tello"
)

// ConsoleRaw console raw flight data
func ConsoleRaw() Manager {
	return &consoleraw{}
}

type consoleraw struct {
	loop bool
}

//

func (consoleraw *consoleraw) Loop(channel <-chan tello.FlightData) {
	consoleraw.loop = true
	for {
		if !consoleraw.loop {
			break
		}
		data := <-channel
		// fieldsMu.Lock()
		updateFields(data)
		log.Printf("%v", fields)
		// fieldsMu.Unlock()
	}
}

func (consoleraw *consoleraw) BreakLoop() {
	consoleraw.loop = false
}
