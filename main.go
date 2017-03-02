package main

import (
	. "./Events"
	. "./Network"
	. "./Network/network/peers"
	. "./driver"
	. "./elevatorController"
	//. ".Network/network/peers"
	"fmt"
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

	go RunNetwork(elevatorData, updateElevatorTxCh, updateElevatorRxCh, newOrderTxCh, newOrderRxCh, peerUpdateCh, peerTxEnableCh)

	go ReadAllSensors2(arriveAtFloorCh, externalButtonCh, internalButtonCh)

	for {
		select {

		case msg1 := <-arriveAtFloorCh:
			//fsmArriveAtFloor(msg)

			elevatorData = FsmArriveAtFloor(elevatorData, msg1)

		case msg2 := <-externalButtonCh:
			//elevatorData = fsmExternalButtonPressed(elevatorData, msg)
			elevatorData = FsmExternalButtonPressed(elevatorData, msg2, newOrderTxCh)

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
			OnlineElevatorsUpdate(msg6)

		}
	}

}

func main1() {
	NetworkTest()

}

func NetworkTest() {

	elevatorData := InitializeElevator()

	updateElevatorRxCh := make(chan ElevatorData, 50)
	updateElevatorTxCh := make(chan ElevatorData, 50)

	newOrderTxCh := make(chan ElevatorOrder, 50)
	newOrderRxCh := make(chan ElevatorOrder, 50)

	peerUpdateCh := make(chan PeerUpdate, 50)
	peerTxEnableCh := make(chan bool)

	arriveAtFloorCh := make(chan int)
	externalButtonCh := make(chan ElevatorOrder, 50)
	internalButtonCh := make(chan int, 50)

	go RunNetwork(elevatorData, updateElevatorTxCh, updateElevatorRxCh, newOrderTxCh, newOrderRxCh, peerUpdateCh, peerTxEnableCh)

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
			fmt.Println(msg2)

		case msg3 := <-internalButtonCh:
			updateElevatorTxCh <- elevatorData
			fmt.Println(msg3)
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
