package main

import (
	"fmt"
	"log"

	"github.com/epaoletta/tello_tests/src/flightdata"
	"github.com/tello"
)

var (
	drone tello.Tello
)

const (
	flightDataUpdatePeriod = 100
)

func main() {
	// app start
	log.Printf("Tello Desktop")
	log.Printf("event: app_start")

	// drone connect control
	log.Printf("event: drone_control_connect")
	err := drone.ControlConnectDefault()
	if err != nil {
		errStr := fmt.Sprintf("drone_control_connect_error: %v", err)
		log.Fatal(errStr)
		panic(errStr)
	}
	defer drone.ControlDisconnect()

	// drone flight data
	log.Printf("event: drone_flight_data")
	flightDataManager, flightDataChannel, err := getFlightDataManager()
	if err != nil {
		errStr := fmt.Sprintf("drone_flight_data_error: %v", err)
		log.Fatal(errStr)
		panic(errStr)
	}
	go flightDataManager.Loop(flightDataChannel)

	/*
		// video feed
		log.Printf("event: drone_video_connect")
		_, err = telloDrone.VideoConnectDefault()
		if err != nil {
			errStr := fmt.Sprintf("drone_video_connect_error: %v", err)
			log.Fatal(errStr)
			panic(errStr)
		}
		defer telloDrone.VideoDisconnect()
		telloDrone.StartVideo()
		_ = gocv.NewWindow("video_feed")


				// video update
			videoBuffer := <-videoChannel
			err := video.Update(videoWindow, videoBuffer)
			if err != nil {
				log.Printf("video_frame_lost_error: %v", err.Error())
			}
	*/

	log.Printf("event: flight_loop")
	var i int64
	for {
		i++

		if i == 99999999 {
			flightDataManager.BreakLoop()
		}
	}

}

//

func getFlightDataManager() (flightdata.Manager, <-chan tello.FlightData, error) {
	manager := flightdata.ConsoleRaw()
	channel, err := drone.StreamFlightData(false, flightDataUpdatePeriod)
	return manager, channel, err
}

// func getVideoManager() (video.Manager, <-chan )
