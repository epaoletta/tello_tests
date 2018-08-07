package flightdata

import (
	"log"

	"github.com/tello"
)

// ConsoleRaw console raw flight data
func ConsoleRaw() Manager {
	return &consoleRaw{}
}

type consoleRaw struct {
	loop bool
}

//

func (consoleRaw *consoleRaw) Loop(channel <-chan tello.FlightData) {
	consoleRaw.loop = true
	for {
		if !consoleRaw.loop {
			break
		}
		data := <-channel
		fieldsMu.Lock()
		updateFields(data)
		log.Printf("%v", fields)
		fieldsMu.Unlock()
	}
}

func (consoleRaw *consoleRaw) BreakLoop() {
	consoleRaw.loop = false
}
