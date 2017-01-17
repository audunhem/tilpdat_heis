package messageHandler

import (
 	"project_elevator/elevator"
	"project_elevator/assigner"
	"project_elevator/externalOrdersWatcher"
	"project_elevator/typeConverter"
	"time"
	"log"
)

func Init(outgoingMessages chan []byte, incomingMessages chan []byte, killChannel chan int, buttonPressed chan []byte, orderComplete chan []byte, updateElevatorData chan []byte, 
	updateOrders chan []byte) {

	id, err := GetID()
	if err != nil {
		log.Println(err)
		log.Println("Elevator will run in offline mode.")
		go StartMessageHandler(id, outgoingMessages, incomingMessages, killChannel, buttonPressed, orderComplete, updateElevatorData, updateOrders)
	} else {
		log.Println("Elevator ID: ", id)
		go StartMessageHandler(id, outgoingMessages, incomingMessages, killChannel, buttonPressed, orderComplete, updateElevatorData, updateOrders)
	}
}

func StartMessageHandler(id byte, outgoingMessages chan []byte, incomingMessages chan []byte, killChannel chan int, buttonPressed chan []byte, orderComplete chan []byte, 
	updateElevatorData chan []byte, updateOrders chan []byte) {

	var online bool = true
	var elevatorData elevator.ElevatorData

	// Channels for Assigner
	killAssigner 		:= make(chan int, 1)
	orderAcknowledged 	:= make(chan []byte, 1024)
	orderQueue 			:= make(chan []byte, 1024)
	costRequest 		:= make(chan []byte, 1024)
	assignOrder 		:= make(chan []byte, 1024)
	costReplies 		:= make(chan []byte, 1024)

	// Channels for ExternalOrdersWatcher (EOW)
	killEOW 			:= make(chan int, 1)
	updateOrderEOW		:= make(chan []byte, 1024)
	costRequestFromEOW	:= make(chan []byte, 1024)

	assigner.Init(id, killAssigner, orderAcknowledged, orderQueue, costRequest, assignOrder, costReplies)
	externalOrdersWatcher.Init(killEOW, updateOrderEOW, costRequestFromEOW)

	for {
		select {
		case message := <-incomingMessages:
			if message[0] == MSG_START && len(message) > 1 {
				receiver := message[1]
				switch receiver {
				case id:
					if len(message) < 6 {break}
					msgType := message[3]

					switch msgType {
					case COST_REPLY:
						if len(message) == 7 {
							ReceiveCostReply(message, costReplies)
						}
						
					case ORDER_ACK:
						if len(message) == 6 {
							ReceiveOrderAcknowledgement(message, orderAcknowledged)
						}

					case ASSIGN_ORDER:
						if len(message) == 6 {
							ReceiveOrderAssignment(message, id, outgoingMessages, updateOrders, updateOrderEOW)
						}
					}

				case REC_ALL:
					if len(message) < 6 {break} 
					msgType := message[3]

					switch msgType {
					case COST_REQUEST:
						if len(message) == 6 {
							ReceiveCostRequest(message, id, elevatorData, outgoingMessages, updateOrderEOW)
						}

					case CLEAR_ORDER:
						if len(message) == 6 {
							sender := message[2]

							select {
							case updateOrderEOW <- []byte{message[4], message[5], 0}:
							default:
								log.Println("UpdateOrderEOW channel is full")
							}

							if sender != id {
								ReceiveClearRequest(message, outgoingMessages, updateOrders)
							}
						}
					}

				case REC_NONE:
					// Ping message, no action to be taken.
				}
			}

		case order := <- orderComplete:
			if online {
				select {
				case outgoingMessages <- MessageOrderComplete(id, order):
				default:
					log.Println("Outgoing messages channel is full")
				}
			} else {
				select {
				case updateOrderEOW <- []byte{order[0], order[1], 0}:
				default:
					log.Println("Update order EOW channel is full")
				}
			}
			
		case elevData := <- updateElevatorData:
			data, err := typeConverter.ConvertMessageToElevatorData(elevData)
			if err != nil {
				log.Println(err)
			} else {
				elevatorData = data
			}

		case order := <- buttonPressed:
			if online {
				select {
				case orderQueue <- order:
				default:
					log.Println("Order queue channel is full")
				}
			} else {
				select {
				case updateOrders <- []byte{order[0], order[1], 1}:
				default:
					log.Println("Update orders channel is full")
				}	
			}

		case costReq := <- costRequest:
			if online {
				select {
				case outgoingMessages <- MessageRequestCost(id, costReq):
				default:
					log.Println("Outgoing messages channel is full")
				}
			} else {
				costReq = []byte{0} // Not necessary when offline
			}

		case order := <- assignOrder:
			if len(order) != 3 {break}
			floor 	:= order[1]
			btnType := order[2]

			if online {
				select {
				case outgoingMessages <- MessageAssignOrder(id, order):
				default:
					log.Println("Outgoing messages channel is full")
				}	
			} else {
				select {
				case updateOrders <- []byte{floor, btnType, 1}:
				default:
					log.Println("Update orders channel is full")
				}	
			}

		case order := <- costRequestFromEOW:
			if len(order) != 2 {break}
			if online {
				select {
				case orderQueue <- order:
				default:
					log.Println("Request queue channel is full")
				}
			} else {
				select {
				case updateOrders <- []byte{order[0], order[1], 1}:
				default:
					log.Println("Update orders channel is full")
				}
			}

		case <-time.After(10 * time.Second): 
			// Check connection status
			select {
			case outgoingMessages <- CreateMessage(REC_NONE, id, 0, []byte{0}):
			default:
				log.Println("Outgoing messages channel is full")
			}
			
		case <- killChannel:
			online = false

			select {
			case killAssigner <- 1:
			default:
				log.Println("Kill Assigner channel is full")
			}

			select {
			case killEOW <- 1:
			default:
				log.Println("Kill EOW channel is full")
			}
		}
	}
}