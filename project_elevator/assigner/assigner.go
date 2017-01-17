package assigner

import (
	"sort"
	"log"
	"project_elevator/timer"
)

type Status int
const (
	Idle Status		= iota
	WaitingForCost	
	WaitingForAck
)

type CurrentSession struct {
	status		Status
	floor 		byte
	buttonType	byte
	attempt		int
}

func Init(id byte, killAssigner chan int, orderAcknowledged chan []byte, orderQueue chan []byte, costRequest chan []byte, assignOrder chan []byte, costReplies chan []byte) {
	go Start(id, killAssigner, orderAcknowledged, orderQueue, costRequest, assignOrder, costReplies)
}

func Start(id byte, killAssigner chan int, orderAcknowledged chan []byte, orderQueue chan []byte, costRequest chan []byte, assignOrder chan []byte, costReplies chan []byte) {

	var session CurrentSession
	session.status = Idle

	var ids []int
	costs := make(map[int]int)

	// Channels for timer routine
	timeOut 	:= make(chan int, 1024)
	startTimer 	:= make(chan int, 1024)
	stopTimer 	:= make(chan int, 1024)

	go timer.InitTimerRoutine(startTimer, timeOut, stopTimer)

	for {
		if session.status == Idle {
			select {
			case order := <- orderQueue:
				costs = ClearMap(costs)
				ids = []int{}
				session.attempt = 0

				select {
				case costRequest <- order:
					session.status = WaitingForCost
				default:
					log.Println("Cost request channel is full")
				}
				
				select {
				case startTimer <- 100:
				default:
					log.Println("Start timer channel is full")
				}
				
				session.floor = order[0]
				session.buttonType = order[1]

			case <- costReplies:
			case <- timeOut: 
			case <- orderAcknowledged:
				// To make sure the channels won't be full
			case <- killAssigner:
				return
			}
		}

		select {
		case costReply := <- costReplies:

			sender 		:= int(costReply[0])
			floor 		:= costReply[1]
			buttonType 	:= costReply[2]
			cost 		:= int(costReply[3])

			if floor == session.floor && buttonType == session.buttonType {
				if cost == 0 {
					session.status = Idle // Cost 0 means that an elevator already has the order
					break
				} else {
					costs[sender] = cost	
				}
			}

		case <- timeOut:
			switch session.status {
			case WaitingForCost:
				session.status = WaitingForAck
				ids = SortKeysByValue(costs)

				if len(ids) > 0 {
					select {
					case assignOrder <- []byte{byte(ids[session.attempt]), session.floor, session.buttonType}:
					default:
						log.Println("Assign order channel is full")
					}
					
					select {
					case startTimer <- 100:
					default:
						log.Println("Start timer channel is full")
					}
					
				} else {
					session.status = Idle
					select {
					case assignOrder <- []byte{id, session.floor, session.buttonType}:
					default:
						log.Println("Assign order channel is full")
					}
				}

			case WaitingForAck:
			 	if session.attempt + 1 < len(ids) {
					session.attempt = session.attempt + 1
					select {
					case assignOrder <- []byte{byte(ids[session.attempt]), session.floor, session.buttonType}:
					default:
						log.Println("Assign order channel is full")
					}
					
					select {
					case startTimer <- 100:
					default:
						log.Println("Start timer channel is full")
					}
					
				} else {
					session.status = Idle
					select {
					case assignOrder <- []byte{id, session.floor, session.buttonType}:
					default:
						log.Println("Assign order channel is full")
					}
				}

			default:
			}

		case order := <- orderAcknowledged:
			floor := order[1]
			buttonType := order[2]

			if floor == session.floor && buttonType == session.buttonType {
				select {
				case stopTimer <- 1:
				default:
					log.Println("Stop timer channel is full")
				}

				session.status = Idle
			}

		case <- killAssigner:
			return

		default: 
			continue
		}
	}
}

func ClearMap(m map[int]int) map[int]int {
	for k := range m {
		delete(m, k)
	}
	return m
}

func SortKeysByValue(m map[int]int) []int {
    n := map[int][]int{}
    
    for k, v := range m {
            n[v] = append(n[v], k)
    }

    var a []int

    for k := range n {
            a = append(a, k)
    }
	
	var keysByValue []int 
	
    sort.Sort(sort.IntSlice(a))
    for _, k := range a {
        for _, s := range n[k] {
			keysByValue = append(keysByValue, s)
        }
    }
	return keysByValue
}