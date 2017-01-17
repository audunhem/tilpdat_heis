package messageHandler

import (
	"log"
	"project_elevator/elevator"
	"project_elevator/fsm"
)

func ReceiveCostReply(message []byte, costReplies chan []byte) {
	sender 		:= message[2]
	floor 		:= message[4]
	buttonType 	:= message[5]
	cost 		:= message[6]
	log.Println("Cost reply from", sender, "for floor:", floor, "and button:", buttonType, "cost:", cost)
	select {
	case costReplies <- []byte{sender, floor, buttonType, cost}:
	default:
		log.Println("Cost replies channel is full")
	}
}

func ReceiveOrderAcknowledgement(message []byte, orderAcknowledged chan []byte) {
	sender 		:= message[2]
	floor 		:= message[4]
	buttonType 	:= message[5]
	log.Println("Received order acknowledgement from", sender)
	select {
	case orderAcknowledged <- []byte{sender, floor, buttonType}:
	default:
		log.Println("OrderAcknowledged channel is full")
	}
}

func ReceiveOrderAssignment(message []byte, id byte, outgoingMessages chan []byte, updateOrders chan []byte, updateOrderEOW chan []byte) {
	sender 		:= message[2]
	floor 		:= message[4]
	buttonType 	:= message[5]
	log.Println("Received order assignment from", sender)

	select {
	case updateOrders <- []byte{floor, buttonType, 1}:
	default:
		log.Println("Update orders channel is full")
	}

	select {
	case updateOrderEOW <- []byte{floor, buttonType, 0}:
	default:
		log.Println("External orders watcher channel is full")
	}
	
	select {
	case outgoingMessages <- MessageAcknowledgeOrder(sender, id, floor, buttonType):
	default:
		log.Println("Outgoing messages channel is full")
	}
}

func ReceiveCostRequest(message []byte, id byte, elevatorData elevator.ElevatorData, outgoingMessages chan []byte, updateOrderEOW chan []byte) {
	sender 		:= message[2]
	floor 		:= message[4]
	buttonType 	:= message[5]
	cost 		:= fsm.FsmOnCostRequest(elevatorData, int(floor), elevator.Button(buttonType))
	log.Println("Cost request received from", sender, "for floor", floor, "and button", buttonType)

	select {
	case outgoingMessages <- MessageCostReply(sender, id, floor, buttonType, cost):
	default:
		log.Println("Outgoing messages channel is full")
	}

	select {
	case updateOrderEOW <- []byte{floor, buttonType, 1}:
	default:
		log.Println("Update order EOW channel is full")
	}	
}

func ReceiveClearRequest(message []byte, outgoingMessages chan []byte, updateOrders chan []byte) {
	floor 		:= message[4]
	buttonType 	:= message[5]
	select {
	case updateOrders <- []byte{floor, buttonType, 0}:
	default:
		log.Println("Update orders channel is full")
	}
}