package main

import (
	"fmt"
	"log"
	"time"

	"github.com/epaoletta/tello_tests/src/flightdata"
	"github.com/epaoletta/tello_tests/src/video"
	"github.com/tello"
)

var (
	drone tello.Tello
)

const (
	showFlightData = true
	showVideo      = true
)

const (
	flightDataUpdatePeriod   = 100
	videoConnectConfirmation = 250
	desiredFPS               = 30
)

func main() {
	// app start
	log.Printf("Tello Desktop")
	log.Printf("event: app_start")

	// drone control connect
	log.Printf("event: drone_control_connect")
	err := drone.ControlConnectDefault()
	if err != nil {
		errStr := fmt.Sprintf("drone_control_connect_error: %v", err)
		log.Fatal(errStr)
		panic(errStr)
	}
	defer drone.ControlDisconnect()

	// drone flight data
	if showFlightData {
		log.Printf("event: drone_flight_data")
		flightDataManager, flightDataChannel, err := getFlightDataManager()
		if err != nil {
			errStr := fmt.Sprintf("drone_flight_data_error: %v", err)
			log.Fatal(errStr)
			panic(errStr)
		}

		go flightDataManager.Loop(flightDataChannel)
	}

	// video feed
	if showVideo {
		log.Printf("event: drone_video_connect")
		drone.StartVideo()
		go func() {
			for {
				drone.StartVideo()
				time.Sleep(videoConnectConfirmation * time.Millisecond)
			}
		}()
		videoManager, videoChannel, err := getVideoManager()
		if err != nil {
			errStr := fmt.Sprintf("drone_video_connect_error: %v", err)
			log.Fatal(errStr)
			panic(errStr)
		}

		go videoManager.Loop(videoChannel)
	}

	log.Printf("event: flight_loop")
	for {

	}

}

//

func getFlightDataManager() (flightdata.Manager, <-chan tello.FlightData, error) {
	manager := flightdata.ConsoleRaw()
	channel, err := drone.StreamFlightData(false, flightDataUpdatePeriod)
	return manager, channel, err
}

func getVideoManager() (video.Manager, <-chan []byte, error) {
	manager := video.ConsoleRaw(desiredFPS)
	channel, err := drone.VideoConnectDefault()
	return manager, channel, err
}
