package flightdata

import (
	"fmt"
	"log"
	"math"

	"github.com/tello"
)

type field struct {
	label string
	value interface{}
}

const (
	fHeight = iota
	fBattery
	fWifiStrength
	fMaxHeight
	fLowBattThresh
	fWifiInterference
	fDerivedSpeed
	fGroundSpeed
	fFwdSpeed
	fLatSpeed
	fVertSpeed
	fBattLow
	fBattCrit
	fBattState
	fGroundVis
	fOvertemp
	fLightStrength
	fOnGround
	fHovering
	fFlying
	fFlyMode
	fCameraState
	fDroneFlyTimeLeft
	fDroneBattLeft
	fVelX
	fVelY
	fVelZ
	fPosX
	fPosY
	fPosZ
	fQatW
	fQatX
	fQatY
	fQatZ
	fTemp
	fRoll
	fPitch
	fYaw
	fHome
	fSSID
	fVersion
	fNumFields
)

var (
	fields [fNumFields]field
)

// LogFlightData log flight data
func LogFlightData(flightData tello.FlightData) {
	updateFields(flightData)
	logFields()
}

func updateFields(flightData tello.FlightData) {
	fields[fHeight].value = fmt.Sprintf("%.1fm", float32(flightData.Height)/10)
	fields[fBattery].value = fmt.Sprintf("%d%%", flightData.BatteryPercentage)
	fields[fWifiStrength].value = fmt.Sprintf("%d%%", flightData.WifiStrength)

	fields[fMaxHeight].value = fmt.Sprintf("%dm", flightData.MaxHeight)
	fields[fLowBattThresh].value = fmt.Sprintf("%d%%", flightData.LowBatteryThreshold)
	fields[fWifiInterference].value = fmt.Sprintf("%d%%", flightData.WifiInterference)

	fields[fDerivedSpeed].value = fmt.Sprintf("%.1fm/s", math.Sqrt(float64(flightData.NorthSpeed*flightData.NorthSpeed)+float64(flightData.EastSpeed*flightData.EastSpeed)))
	fields[fGroundSpeed].value = fmt.Sprintf("%dm/s", flightData.GroundSpeed)
	fields[fFwdSpeed].value = fmt.Sprintf("%dm/s", flightData.NorthSpeed)
	fields[fLatSpeed].value = fmt.Sprintf("%dm/s", flightData.EastSpeed)

	fields[fVertSpeed].value = fmt.Sprintf("%dm/s", flightData.VerticalSpeed)

	fields[fBattLow].value = boolToYN(flightData.BatteryLow)
	fields[fBattCrit].value = boolToYN(flightData.BatteryCritical)
	fields[fBattState].value = boolToYN(flightData.BatteryState)

	fields[fGroundVis].value = boolToYN(flightData.DownVisualState)
	fields[fOvertemp].value = boolToYN(flightData.OverTemp)
	fields[fLightStrength].value = fmt.Sprintf("%d", flightData.LightStrength)

	fields[fOnGround].value = boolToYN(flightData.OnGround)
	fields[fHovering].value = boolToYN(flightData.DroneHover)
	fields[fFlying].value = boolToYN(flightData.Flying)

	fields[fFlyMode].value = fmt.Sprintf("%d", flightData.FlyMode)

	fields[fCameraState].value = fmt.Sprintf("%d", flightData.CameraState)
	fields[fDroneFlyTimeLeft].value = fmt.Sprintf("%d", flightData.DroneFlyTimeLeft)
	fields[fDroneBattLeft].value = fmt.Sprintf("%dmV", flightData.BatteryMilliVolts)

	fields[fVelX].value = fmt.Sprintf("%dcm/s", flightData.MVO.VelocityX)
	fields[fVelY].value = fmt.Sprintf("%dcm/s", flightData.MVO.VelocityY)
	fields[fVelZ].value = fmt.Sprintf("%dcm/s", flightData.MVO.VelocityZ)

	fields[fPosX].value = fmt.Sprintf("%f", flightData.MVO.PositionX)
	fields[fPosY].value = fmt.Sprintf("%f", flightData.MVO.PositionY)
	fields[fPosZ].value = fmt.Sprintf("%f", flightData.MVO.PositionZ)

	fields[fQatW].value = fmt.Sprintf("%f", flightData.IMU.QuaternionW)
	fields[fQatX].value = fmt.Sprintf("%f", flightData.IMU.QuaternionX)
	fields[fQatY].value = fmt.Sprintf("%f", flightData.IMU.QuaternionY)
	fields[fQatZ].value = fmt.Sprintf("%f", flightData.IMU.QuaternionZ)
	fields[fTemp].value = fmt.Sprintf("%dC", flightData.IMU.Temperature)

	// p, r, y := tello.QuatToEulerDeg(flightData.IMU.QuaternionX, flightData.IMU.QuaternionY, flightData.IMU.QuaternionZ, flightData.IMU.QuaternionW)
	// fields[fRoll].value = fmt.Sprintf("%d", r)
	// fields[fPitch].value = fmt.Sprintf("%d", p)
	fields[fYaw].value = fmt.Sprintf("%d°", flightData.IMU.Yaw)

	fields[fSSID].value = flightData.SSID
	fields[fVersion].value = flightData.Version
}

func boolToYN(b bool) string {
	if b {
		return "Y"
	}
	return "N"
}

func logFields() {
	log.Printf("%v", fields)
}
