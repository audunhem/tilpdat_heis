package timer

import "time"

var startTime int64
var timerActive bool

func TimerStart() {
	timerActive = true
	startTime = time.Now().Unix()
}

func TimerStop() {
	timerActive = false
}

func TimerTimeout() bool {
	if time.Now().Unix() - startTime >= 3 && timerActive {
		return true
	}
	return false
}