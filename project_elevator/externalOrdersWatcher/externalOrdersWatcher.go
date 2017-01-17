package externalOrdersWatcher

import (
	"time"
	"log"
	"project_elevator/elevator"
)

const tresholdTime = 10 // Seconds

func Init(killChannel chan int, newOrderChannel chan []byte, requestCostChannel chan []byte) {
	go Start(killChannel, newOrderChannel, requestCostChannel)
}

func Start(killChannel chan int, newOrderChannel chan []byte, requestCostChannel chan []byte) {

	var orderSetTime [elevator.N_FLOORS][2]int64

	for {
		select{
		case <- killChannel:
			log.Println("I know I'm running offline.")

		case order := <- newOrderChannel:
			floor 		:= order[0]
			buttonType 	:= order[1]
			value		:= order[2]
			
			if floor < elevator.N_FLOORS || buttonType < 2 {
				if value == 1 {
					orderSetTime[floor][buttonType] = time.Now().Unix()
				} else {
					orderSetTime[floor][buttonType] = 0
				}
			} else {
				log.Println("Invalid floor or buttonType in External Orders Watcher")
			}

		default:
			for f := range(orderSetTime) {
				for b := range(orderSetTime[f]) {
					if orderSetTime[f][b] == 0 {
						continue
					} else {
						if time.Now().Unix() - orderSetTime[f][b] > tresholdTime {
							select {
							case requestCostChannel <- []byte{byte(f), byte(b)}:
								orderSetTime[f][b] = time.Now().Unix()
							default:
								log.Println("Request cost channel is full")
							}		
						}
					}
				}
			}
		}
	}	
}