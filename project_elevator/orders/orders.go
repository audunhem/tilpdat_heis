package orders

import (
	"project_elevator/elevator"
	"project_elevator/fileManager"
	"log"
)

const ( 
	D_Down elevator.Dirn = elevator.D_Down
	D_Stop = elevator.D_Stop
	D_Up = elevator.D_Up
)

const (
	B_HallUp = elevator.B_HallUp
	B_HallDown = elevator.B_HallDown
	B_Cab = elevator.B_Cab
)

const N_FLOORS = elevator.N_FLOORS
const N_BUTTONS = elevator.N_BUTTONS

func OrdersAbove(elevatorData elevator.ElevatorData) bool {
	for f := elevatorData.Floor+1; f < N_FLOORS; f++{
		for btn := 0; btn < N_BUTTONS; btn++{
			if elevatorData.Orders[f][btn] == 1{
				return true
			}
		}
	}
	return false
} 

func OrdersBelow(elevatorData elevator.ElevatorData) bool {
	for f := 0; f < elevatorData.Floor; f++{
		for btn := 0; btn < N_BUTTONS; btn++{
			if elevatorData.Orders[f][btn] == 1{
				return true
			}
		}
	}
	return false
}

func OrdersChooseDirection(elevatorData elevator.ElevatorData) elevator.Dirn {
	switch elevatorData.Direction{
	case D_Up:
		if OrdersAbove(elevatorData){
			return D_Up
		} else if OrdersBelow(elevatorData){
			return D_Down
		} else{
			return D_Stop
		}
	case D_Down:
		if OrdersBelow(elevatorData){
			return D_Down
		} else if OrdersAbove(elevatorData){
			return D_Up
		} else{
			return D_Stop
		}
	case D_Stop:
		if OrdersBelow(elevatorData){
			return D_Down
		} else if OrdersAbove(elevatorData){
			return D_Up
		} else{
			return D_Stop
		}
	default:
		return D_Stop
	}
}

func OrdersShouldStop(elevatorData elevator.ElevatorData) bool {
	switch elevatorData.Direction{
	case D_Down:
		if elevatorData.Orders[elevatorData.Floor][B_HallDown] == 1 ||
		 elevatorData.Orders[elevatorData.Floor][B_Cab] == 1 || !OrdersBelow(elevatorData){
		 	return true
		}
	case D_Up:
		if elevatorData.Orders[elevatorData.Floor][B_HallUp] == 1 ||
		 elevatorData.Orders[elevatorData.Floor][B_Cab] == 1 || !OrdersAbove(elevatorData){
		 	return true
		 }
	case D_Stop:
	default:
		return true
	}
	return false
}

func OrdersClearAtCurrentFloor(elevatorData elevator.ElevatorData, orderComplete chan []byte) elevator.ElevatorData {
	elevatorData.Orders[elevatorData.Floor][B_Cab] = 0;
	elevator.ElevatorWriteOrderButtonLight(elevatorData.Floor, B_Cab, 0)
	switch elevatorData.Direction{
	case D_Up:
		elevatorData.Orders[elevatorData.Floor][B_HallUp] = 0;
		elevator.ElevatorWriteOrderButtonLight(elevatorData.Floor, B_HallUp, 0)
		select {
		case orderComplete <- []byte{byte(elevatorData.Floor), byte(B_HallUp)}:
		default:
			log.Println("Order complete channel is full")
		}

		if !OrdersAbove(elevatorData){
			elevatorData.Orders[elevatorData.Floor][B_HallDown] = 0
			elevator.ElevatorWriteOrderButtonLight(elevatorData.Floor, B_HallDown, 0)
			select {
			case orderComplete <- []byte{byte(elevatorData.Floor), byte(B_HallDown)}:
			default:
				log.Println("Order complete channel is full")
			}
		}

	case D_Down:
		elevatorData.Orders[elevatorData.Floor][B_HallDown] = 0
		elevator.ElevatorWriteOrderButtonLight(elevatorData.Floor, B_HallDown, 0)
		select {
		case orderComplete <- []byte{byte(elevatorData.Floor), byte(B_HallDown)}:
		default:
			log.Println("Order complete channel is full")
		}
		
		if !OrdersBelow(elevatorData){
			elevatorData.Orders[elevatorData.Floor][B_HallUp] = 0
			elevator.ElevatorWriteOrderButtonLight(elevatorData.Floor, B_HallUp, 0)
			select {
			case orderComplete <- []byte{byte(elevatorData.Floor), byte(B_HallUp)}:
			default:
				log.Println("Order complete channel is full")
			}
		}		
	
	default:
		elevatorData.Orders[elevatorData.Floor][B_HallUp] = 0
		elevatorData.Orders[elevatorData.Floor][B_HallDown] = 0
		elevator.ElevatorWriteOrderButtonLight(elevatorData.Floor, B_HallUp, 0)
		elevator.ElevatorWriteOrderButtonLight(elevatorData.Floor, B_HallDown, 0)
		select {
		case orderComplete <- []byte{byte(elevatorData.Floor), byte(B_HallUp)}:
		default:
			log.Println("Order complete channel is full")
		}
		select {
		case orderComplete <- []byte{byte(elevatorData.Floor), byte(B_HallDown)}:
		default:
			log.Println("Order complete channel is full")
		}
		
	}
	
	fileManager.WriteOrdersToFile(elevatorData)
	return elevatorData
}

func OrdersCostGoingUp(elevatorData elevator.ElevatorData, orderFloor int, orderButtonType elevator.Button) int {
	cost := 1
	switch orderButtonType{
	case B_HallUp:
		if elevatorData.Floor >= orderFloor {
			for i := elevatorData.Floor; i < N_FLOORS; i++{
				if elevatorData.Orders[i][B_HallUp] == 1 {
					cost += 2
				}
			}
			for i := 0; i < orderFloor; i++ {
				if elevatorData.Orders[i][B_HallUp] == 1 {
					cost += 2
				}
			}
			for i := 0; i < N_FLOORS; i++ {
				if elevatorData.Orders[i][B_HallDown] == 1 || elevatorData.Orders[i][B_Cab] == 1 {
					cost += 2
				}
			}
		} else {
			for i := elevatorData.Floor; i < orderFloor; i++ {
				if elevatorData.Orders[i][B_HallUp] == 1 || elevatorData.Orders[i][B_Cab] == 1 {
					cost += 2
				}
			}

		}
	case B_HallDown:
		if elevatorData.Floor < orderFloor {
			for i := elevatorData.Floor; i < N_FLOORS; i++ {
				if elevatorData.Orders[i][B_HallUp] == 1 || elevatorData.Orders[i][B_Cab] == 1 {
					cost += 2
				}
			}
			for i := N_FLOORS-1; i > orderFloor; i-- {
				if elevatorData.Orders[i][B_HallUp] == 1 {
					cost += 2
				}
			}
		}else {
			for i := elevatorData.Floor; i < N_FLOORS; i++ {
				if elevatorData.Orders[i][B_HallDown] == 1 {
					cost += 2
				}
			}
			for i := N_FLOORS-1; i > orderFloor; i-- {
				if elevatorData.Orders[i][B_HallUp] == 1 || elevatorData.Orders[i][B_Cab] == 1 {
					cost += 2
				}
			}
		}

	}
	return cost
}

func OrdersCostGoingDown(elevatorData elevator.ElevatorData, orderFloor int, orderButtonType elevator.Button) int{
	cost := 1
	switch orderButtonType {
	case B_HallDown:
		if elevatorData.Floor <= orderFloor {
			for i := elevatorData.Floor; i >= 0; i-- {
				if elevatorData.Orders[i][B_HallDown] == 1 {
					cost += 2
				}
			}
			for i := 0; i < N_FLOORS; i++ {
				if elevatorData.Orders[i][B_HallUp] == 1 || elevatorData.Orders[i][B_Cab] == 1 {
					cost += 2
				}
			}
			for i := N_FLOORS-1; i > orderFloor; i-- {
				if elevatorData.Orders[i][B_HallDown] == 1 {
					cost += 2
				}
			}
		}else if elevatorData.Floor > orderFloor {
			for i := elevatorData.Floor; i > orderFloor; i-- {
				if elevatorData.Orders[i][B_HallDown] == 1 || elevatorData.Orders[i][B_Cab] == 1 {
					cost += 2
				}
			}
		}
	case B_HallUp:
		if elevatorData.Floor < orderFloor {
			for i := elevatorData.Floor; i >= 0; i-- {
				if elevatorData.Orders[i][B_HallDown] == 1 {
					cost += 2
				}
			} 
			for i := 0; i < orderFloor; i++ {
				if elevatorData.Orders[i][B_HallUp] == 1 || elevatorData.Orders[i][B_Cab] == 1 {
					cost += 2
				}
			}
		} else {
			for i := elevatorData.Floor; i >= 0; i-- {
				if elevatorData.Orders[i][B_HallDown] == 1 || elevatorData.Orders[i][B_Cab] == 1 {
					cost += 2
				}
			}
			for i := 0; i < orderFloor; i++ {
				if elevatorData.Orders[i][B_HallUp] == 1 {
					cost += 2
				}
			}
		}
	}
	return cost
}

func OrdersCostByDistance(elevatorData elevator.ElevatorData, orderFloor int) int{
	cost := 1
	if elevatorData.Floor > orderFloor{
		cost += (elevatorData.Floor - orderFloor)
	} else {
		cost += (orderFloor - elevatorData.Floor)
	}
	return cost
}