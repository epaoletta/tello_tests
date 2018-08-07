package video

import "gocv.io/x/gocv"

// Update updates the video window
func Update(videoWindow *gocv.Window, videoBuffer []byte) error {
	gocv.NewMat()
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
