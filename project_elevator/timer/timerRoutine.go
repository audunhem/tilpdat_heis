package timer

import (
	"time"
	"log"
)

func InitTimerRoutine(startTimer chan int, timeout chan int, stopTimer chan int) {
	var running bool
	for {
		select {
		case start := <- startTimer:
			if !running && start != 0{
				running = true
				select {
				case stop := <- stopTimer:
					if running && stop == 1 {
						running = false
					}
				case <-time.After(time.Duration(start) * time.Millisecond):
					if running{
						select {
						case timeout <- 1:
						default:
							log.Println("Timeout channel is full")						
						}
						running = false
					}
				}
			}

		default:
		}
	}
}