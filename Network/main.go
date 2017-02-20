package main

import (
	."../driver/"
)

func main() {


	updateTx = make(chan ElevatorData)
	updateRx = make(chan ElevatorData)

	newOrderTx = make(chan ElevatorOrder)
	newOrderRx = make(chan ElevatorOrder)

	peerUpdateCh = make(chan peers.PeerUpdate)

	runNetwork(updateTx, updateRx, newOrderTx, newOrderRx, peerUpdateCh)

	for {}


}
