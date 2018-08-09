package video

import (
	"log"
	"time"
)

// ConsoleRaw console raw video manager
func ConsoleRaw(desiredFPS int) Manager {
	return &consoleraw{
		desiredFPS: desiredFPS,
	}
}

type consoleraw struct {
	loop bool

	desiredFPS int
}

//

func (consoleraw *consoleraw) Loop(channel <-chan []byte) {
	consoleraw.loop = true
	frame := uint64(0)
	for {
		if !consoleraw.loop {
			break
		}
		_ = <-channel
		frame++
		// videoMu.Lock()
		log.Printf("frame: %d", frame)
		//log.Printf("%v", videoBuffer)
		// videoMu.Unlock()
		time.Sleep(30 * time.Millisecond)
	}

}

func (consoleraw *consoleraw) BreakLoop() {
	consoleraw.loop = false
}
