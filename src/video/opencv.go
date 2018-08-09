package video

import (
	"log"
	"time"

	"gocv.io/x/gocv"
)

// OpenCV opencv video manager
func OpenCV(desiredFPS int) Manager {
	return &openCV{
		desiredFPS: desiredFPS,
	}
}

type openCV struct {
	loop bool

	desiredFPS int
}

//

func (opencv *openCV) Loop(channel <-chan []byte) {
	opencv.loop = true
	videoWindow := gocv.NewWindow("video_feed")
	time.Sleep(time.Second)
	for {
		if !opencv.loop {
			break
		}
		videoBuffer := <-channel
		videoMu.Lock()
		if err := Update(videoWindow, videoBuffer); err != nil {
			log.Printf("video_frame_lost_error: %v", err.Error())
		}
		videoMu.Unlock()
	}

}

func (opencv *openCV) BreakLoop() {
	opencv.loop = false
}

//

// Update updates the video window
func Update(videoWindow *gocv.Window, videoBuffer []byte) error {
	img, err := gocv.NewMatFromBytes(720, 960, gocv.MatTypeCV8UC3, videoBuffer)
	if err != nil {
		return err
	}
	if img.Empty() {
		return nil
	}
	videoWindow.IMShow(img)
	videoWindow.WaitKey(1)
	return nil
}
