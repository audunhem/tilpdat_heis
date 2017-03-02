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


func NetworkTest() {

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
		var testOrder ElevatorOrder
		testOrder.Floor = 0
		testOrder.Direction = 0
		testOrder.ElevatorID = "test"

		newOrderTxCh <- testOrder

	case msg3 := <-internalButtonCh:
		updateElevatorTxCh <-elevatorData

	case msg4 := <-updateElevatorRxCh:
		fmt.Println("Elevator Update received")
		fmt.Println(msg4)
		//elevatorData = OrderReceivedUpdate(elevatorData, msg)

	case msg5 := <-newOrderRxCh:
		fmt.Println("New order received")
		fmt.Println(msg5)
		//elevatorData = OrderReceivedOrder(elevatorData, msg)
	case msg6 := <-peerUpdateCh:
		fmt.Println(msg6)
		//elevatorData = PeerUpdate(elevatorData, msg)

	}
}

}
