package main

import (
	. "./driver"
	. "./elevatorController"
	"fmt"
	"time"
)

func main() {
	updatedDataFSM := make(chan ElevatorData)
	currentFloorChannel := make(chan int)
	newOrderButtonTypeChannel := make(chan ButtonType)
	newOrderFloorChannel := make(chan int)

	go ReadAllSensors(updatedDataFSM, currentFloorChannel, newOrderButtonTypeChannel, newOrderFloorChannel)
	go print(currentFloorChannel)
	fmt.Println("ok")
	GoToFloor(1)
	GoToFloor(2)
	GoToFloor(3)
}

func print(currentFloorChannel chan int) {
	for {
		msg := <-currentFloorChannel
		Println(msg)
		time.Sleep(1 * time.Second)
	}
}

/*InitElevator()
fmt.Println("Press STOP button to stop elevator and exit program.")

GoToFloor(2)

fmt.Println(GetMotorDirection())
GoToFloor(1)

for{
	if GetStopSignal() == 1 {
		SetMotorDirection(DirnStop)

	}
}*/

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
