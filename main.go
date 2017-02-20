package main

import (
	. "./driver"
	. "./elevatorController"
	"fmt"
	"time"
)

func main() {

	elevatorData := InitializeElevator()

	arriveAtFloorCh := make(chan ElevatorOrder)
	externalButtonCh := make(chan ElevatorOrder)
	internalButtonCh := make(chan ElevatorOrder)

	updateElevatorRxCh = make(chan elevatorData, 1024)
	updateElevatorTxCh = make(chan elevatorData)

	newOrderRxCh = make(chan ElevatorOrder)
	newOrderTxCh = make(chan ElevatorOrder)

	peerUpdateCh = make(chan peers)

	go ReadAllSensors2(arriveAtFloorCh, externalButtonCh, internalButtonCh)

	for {
		select {

		case msg <- arriveAtFloorCh
			fsmArriveAtFloor(msg)

		case msg <- externalButtonCh
			elevatorData = fsmExternalButtonPressed(elevatorData, msg)

		case msg <- internalButtonCh
			elevatorData = fsmInternalButtonPressed(elevatorData, msg)

		case msg <- updateElevatorRxCh
			elevatorData = OrderReceivedUpdate(elevatorData, msg)

		case msg <- newOrderRxCh
			elevatorData = OrderReceivedOrder(elevatorData, msg)

		case msg <- peerUpdateCh


	}
	}


}

/*

func main2() {
	updatedDataFSM := make(chan ElevatorData)
	currentFloorChannel := make(chan int)
	newOrderButtonTypeChannel := make(chan ButtonType)
	newOrderFloorChannel := make(chan int)

	var updatedData ElevatorData

	var updatedDataPtr *ElevatorData

	updatedDataPtr = &updatedData
	previousData := InitializeElevator()
	go ReadAllSensors(previousData, updatedDataFSM, currentFloorChannel, newOrderButtonTypeChannel, newOrderFloorChannel)
	go updateDataFromSensor(updatedDataFSM, updatedDataPtr)
	go print(currentFloorChannel)
	fmt.Println("testmain")

	GoToFloor(1, updatedDataPtr)

	GoToFloor(3, updatedDataPtr)
}

func print(currentFloorChannel chan int) {

	for {
		select {
		case msg1 := <-currentFloorChannel:
			fmt.Println(msg1)

		default:
			time.Sleep(1 * time.Second)
		}

	}
}

func updateDataFromSensor(updatedDataFSM chan ElevatorData, updatedData *ElevatorData) {

	for {
		select {

		case update := <-updatedDataFSM:

			(*updatedData) = update
		}

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
