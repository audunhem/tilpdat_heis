package main

import (
	"../driver/elevator_definitions"
)
type HelloMsg struct {
	Message string
	Iter    int
}

func main() {


	updateTx = make(chan driverElevatorData)
	updateRx = make(chan ElevatorData)

	newOrderTx = make(chan ElevatorData)
	newOrderRx = make(chan ElevatorData)

	peerUpdateCh = make(chan peers.PeerUpdate)

	runNetwork(updateTx, updateRx, newOrderTx, newOrderRx, peerUpdateCh)

	for {}


}
