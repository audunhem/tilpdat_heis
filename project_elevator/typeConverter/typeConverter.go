package typeConverter

import (
	"project_elevator/elevator"
	"errors"
)
const N_FLOORS = elevator.N_FLOORS
const N_BUTTONS = elevator.N_BUTTONS

func ConvertElevatorDataToMessage(currentData elevator.ElevatorData) []byte {
	msgSize := (N_FLOORS * N_BUTTONS) + 4
	msg := make([]byte, msgSize)
	msg[0] = byte(currentData.Floor)
	msg[1] = byte(currentData.Direction)

	index := 2
	for f:= 0; f < N_FLOORS; f++ {
		for b:= 0; b < N_BUTTONS; b++ {
			msg[index] = byte(currentData.Orders[f][b])
			index++
		}
	}

	msg[index] = byte(currentData.Behaviour)
	index++
	msg[index] = byte(currentData.Initiated)
	return msg
}

func ConvertMessageToElevatorData(msg []byte) (elevator.ElevatorData, error) {
	
	if len(msg) >= (N_FLOORS * N_BUTTONS) + 4 {
		var orders [N_FLOORS][N_BUTTONS]int
		index := 2

		for f:= 0; f < N_FLOORS; f++ {
			for b:= 0; b < N_BUTTONS; b++ {
				orders[f][b] = int(msg[index])
				index++
			}
		}

		currentData := elevator.ElevatorData {
			Floor:		int(msg[0]),
			Direction:	elevator.Dirn(msg[1]),
			Orders:		orders,
			Behaviour:	elevator.ElevatorBehaviour(msg[index]),
			Initiated:	int(msg[index+1]),
		}
		return currentData, nil
	}
	return elevator.ElevatorData{}, errors.New("Invalid message size in ConvertMessageToElevatorData")
}