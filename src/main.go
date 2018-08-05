package main

import (
	"fmt"
	"log"
	"time"

	"github.com/epaoletta/tello_tests/src/flightdata"
	"github.com/tello"
)

var (
	drone tello.Tello
)

func main() {
	// app start
	log.Printf("Tello Desktop")
	log.Printf("event: app_start")

	// connect to drone
	log.Printf("event: connect_to_drone")
	err := drone.ControlConnectDefault()
	if err != nil {
		errStr := fmt.Sprintf("connection_error: %v", err)
		log.Fatal(errStr)
		panic(errStr)
	}

	// log flight data
	log.Printf("event: flight_data")
	go flightDataLoop()

	log.Printf("event: flight_loop")
	for {

	}

}

func flightDataLoop() {
	for {
		flightData := drone.GetFlightData()
		flightdata.LogFlightData(flightData)

		time.Sleep(50 * time.Millisecond)
	}
}
