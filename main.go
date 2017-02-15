package main

import (
	. "./driver"
	"fmt"
	."./elevatorController"
)

func main() {


	InitElevator()
	fmt.Println("Press STOP button to stop elevator and exit program.")

	GoToFloor(2)

	fmt.Println(GetMotorDirection())
	GoToFloor(1)

	for{
		if GetStopSignal() == 1 {
			SetMotorDirection(DirnStop)

		}
	}


	/*SetMotorDirection(DirnUp)

	SetFloorIndicator(3)
	for {
		GetButtonSignal(0, 2)
		if GetFloorSensorSignal() == N_FLOORS-1 {
			SetMotorDirection(DirnDown)

		} else if GetFloorSensorSignal() == 0 {
			SetMotorDirection(DirnUp)
		}

		if GetFloorSensorSignal() == N_FLOORS-1 {
		}
		// Stop elevator and exit program if the stop button is pressed
		if GetStopSignal() == 1 {
			SetMotorDirection(DirnStop)
		}
	}

	*/


}
