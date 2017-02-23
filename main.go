package main

import (
	. "./Network"
	. "./Network/network/peers"
	. "./driver"
	//. "./elevatorController"
	//. ".Network/network/peers"
	//"fmt"
	/*"time" */)

func main() {
	elevatorData := InitializeElevator()

	updateElevatorRxCh := make(chan ElevatorData)
	updateElevatorTxCh := make(chan ElevatorData)

	newOrderTxCh := make(chan ElevatorOrder)
	newOrderRxCh := make(chan ElevatorOrder)

	peerUpdateCh := make(chan PeerUpdate)
	peerTxEnableCh := make(chan bool)

	arriveAtFloorCh := make(chan int)
	externalButtonCh := make(chan ElevatorOrder, 10)
	internalButtonCh := make(chan int, 10)

	go RunNetwork(updateElevatorTxCh, updateElevatorRxCh, newOrderTxCh, newOrderRxCh, peerUpdateCh, peerTxEnableCh)

	go ReadAllSensors2(arriveAtFloorCh, externalButtonCh, internalButtonCh)

	for {
		select {

		case msg1 := <-arriveAtFloorCh:
			//fsmArriveAtFloor(msg)

			elevatorData = FsmArriveAtFloor(elevatorData, msg1)

		case msg2 := <-externalButtonCh:
			//elevatorData = fsmExternalButtonPressed(elevatorData, msg)
			elevatorData = FsmExternalButtonPressed(elevatorData, msg2)

			PrintOrderList(elevatorData)

		case msg3 := <-internalButtonCh:
			elevatorData = FsmInternalButtonPressed(elevatorData, msg3)
			PrintOrderList(elevatorData)

		case msg4 := <-updateElevatorRxCh:
			fmt.Println(msg4)
			//elevatorData = OrderReceivedUpdate(elevatorData, msg)

		case msg5 := <-newOrderRxCh:
			fmt.Println(msg5)
			//elevatorData = OrderReceivedOrder(elevatorData, msg)
		case msg6 := <-peerUpdateCh:
			fmt.Println(msg6)
			//elevatorData = PeerUpdate(elevatorData, msg)

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
