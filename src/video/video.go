package video

import "sync"

// Manager video manager interface
type Manager interface {
	Loop(<-chan []byte)
	BreakLoop()
}

//

var (
	videoMu sync.RWMutex
)
