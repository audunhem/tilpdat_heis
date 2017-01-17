package fsm

import (
	"project_elevator/elevator"
	"project_elevator/orders"
	"project_elevator/timer"
	"project_elevator/fileManager"
	"log"
)

const N_FLOORS = elevator.N_FLOORS
const N_BUTTONS = elevator.N_BUTTONS

func SetAllLights(currentData elevator.ElevatorData) {
	for floor := 0; floor < N_FLOORS; floor++{
		for btn := elevator.Button(0); btn < N_BUTTONS; btn++{
			elevator.ElevatorWriteOrderButtonLight(floor, btn, currentData.Orders[floor][btn])
		}
	}
}

func FsmOnInitBetweenFloors(elevatorData elevator.ElevatorData) elevator.ElevatorData {
	if elevatorData.Initiated == 0 {
		elevator.ElevatorWriteMotorDirection(elevator.D_Down)
		elevatorData.Direction = elevator.D_Down
		elevatorData.Behaviour = elevator.EB_Moving
	}
	return elevatorData
}

func FsmOnOrderButtonPressed(btn_floor int, btn_type elevator.Button, elevatorData elevator.ElevatorData, buttonPressed chan []byte, orderComplete chan []byte) elevator.ElevatorData {

	switch elevatorData.Behaviour{
	case elevator.EB_DoorOpen:
		if elevatorData.Floor == btn_floor{
			timer.TimerStart()
		} else {
			elevator.ElevatorWriteOrderButtonLight(btn_floor, btn_type, 1)
			if btn_type != elevator.B_Cab {
				select {
				case buttonPressed <- []byte{byte(btn_floor), byte(btn_type)}:
				default:
					log.Println("Button pressed channel is full")
				}
			} else {
				elevatorData.Orders[btn_floor][btn_type] = 1;
			}	
		}

	case elevator.EB_Moving:
		elevator.ElevatorWriteOrderButtonLight(btn_floor, btn_type, 1)
		if btn_type != elevator.B_Cab {
			select {
			case buttonPressed <- []byte{byte(btn_floor), byte(btn_type)}:
			default:
				log.Println("Button pressed channel is full")
			}
		} else {
			elevatorData.Orders[btn_floor][btn_type] = 1;
		}	 

	case elevator.EB_Idle:
		if elevatorData.Floor == btn_floor {
			elevator.ElevatorWriteDoorLight(1)
			timer.TimerStart()
			elevatorData = orders.OrdersClearAtCurrentFloor(elevatorData, orderComplete)
			elevatorData.Behaviour = elevator.EB_DoorOpen
		} else {
			elevator.ElevatorWriteOrderButtonLight(btn_floor, btn_type, 1)
			if btn_type != elevator.B_Cab {
				select{
				case buttonPressed <- []byte{byte(btn_floor), byte(btn_type)}:
				default:
					log.Println("Button pressed channel is full")
				}
			} else {
				elevatorData.Orders[btn_floor][btn_type] = 1;
				elevatorData.Direction = orders.OrdersChooseDirection(elevatorData)
				elevator.ElevatorWriteMotorDirection(elevatorData.Direction)
				elevatorData.Behaviour = elevator.EB_Moving
			}
		}
	}

	fileManager.WriteOrdersToFile(elevatorData)
	return elevatorData
}

func FsmOnDoorTimeout(elevatorData elevator.ElevatorData) elevator.ElevatorData {

	switch elevatorData.Behaviour {
	case elevator.EB_DoorOpen :
		elevatorData.Direction = orders.OrdersChooseDirection(elevatorData)

		elevator.ElevatorWriteDoorLight(0)
		elevator.ElevatorWriteMotorDirection(elevatorData.Direction)

		if elevatorData.Direction == elevator.D_Stop{
			elevatorData.Behaviour = elevator.EB_Idle
		} else {
			elevatorData.Behaviour = elevator.EB_Moving
		}
	default :
	}

	return elevatorData
}

func FsmOnFloorArrival(newFloor int, elevatorData elevator.ElevatorData, orderComplete chan []byte) elevator.ElevatorData {

	elevatorData.Floor = newFloor

	elevator.ElevatorWriteFloorIndicator(elevatorData.Floor)

	switch elevatorData.Behaviour {
	case elevator.EB_Moving:
		if orders.OrdersShouldStop(elevatorData) {
			elevator.ElevatorWriteMotorDirection(elevator.D_Stop)
			elevator.ElevatorWriteDoorLight(1)
			elevatorData = orders.OrdersClearAtCurrentFloor(elevatorData, orderComplete)
			timer.TimerStart()
			elevatorData.Behaviour = elevator.EB_DoorOpen
		}
	case elevator.EB_Idle:
		if orders.OrdersShouldStop(elevatorData) {
			elevator.ElevatorWriteDoorLight(1)
			elevatorData = orders.OrdersClearAtCurrentFloor(elevatorData, orderComplete)
			timer.TimerStart()
			elevatorData.Behaviour = elevator.EB_DoorOpen
		} else {
			elevatorData.Direction = orders.OrdersChooseDirection(elevatorData)
			if elevatorData.Direction!= elevator.D_Stop{
				elevator.ElevatorWriteMotorDirection(elevatorData.Direction)
				elevatorData.Behaviour = elevator.EB_Moving
			}
		}
	default:

	}

	return elevatorData
}

func FsmOnCostRequest(elevatorData elevator.ElevatorData, orderFloor int, orderButtonType elevator.Button) int {
	elevator.ElevatorWriteOrderButtonLight(orderFloor, orderButtonType, 1)
	cost := 0
	if elevatorData.Orders[orderFloor][orderButtonType] == 1 {
		return cost
	}
	switch elevatorData.Behaviour{
	case elevator.EB_Moving:
		switch elevatorData.Direction{
		case elevator.D_Up:
			cost = orders.OrdersCostByDistance(elevatorData, orderFloor) + orders.OrdersCostGoingUp(elevatorData, orderFloor, orderButtonType)
		case elevator.D_Down:
			cost = orders.OrdersCostByDistance(elevatorData, orderFloor) + orders.OrdersCostGoingDown(elevatorData, orderFloor, orderButtonType)
		case elevator.D_Stop:
			cost = orders.OrdersCostByDistance(elevatorData, orderFloor)
		}
	case elevator.EB_DoorOpen:
		cost = orders.OrdersCostByDistance(elevatorData, orderFloor)
	case elevator.EB_Idle:
		cost = orders.OrdersCostByDistance(elevatorData, orderFloor)
	}
	return cost
}

func FsmOnOrdersUpdate(elevatorData elevator.ElevatorData, orderComplete chan []byte) elevator.ElevatorData{
	
	switch elevatorData.Behaviour{
	case elevator.EB_Idle:
		elevatorData.Direction = orders.OrdersChooseDirection(elevatorData)

		if elevatorData.Direction == elevator.D_Stop{
			elevator.ElevatorWriteDoorLight(1)
			timer.TimerStart()
			elevatorData = orders.OrdersClearAtCurrentFloor(elevatorData, orderComplete)
			elevatorData.Behaviour = elevator.EB_DoorOpen
		} else {
			elevator.ElevatorWriteMotorDirection(elevatorData.Direction)
			elevatorData.Behaviour = elevator.EB_Moving
		}

	case elevator.EB_DoorOpen:
		timer.TimerStart()
		elevatorData = orders.OrdersClearAtCurrentFloor(elevatorData, orderComplete)
	}

	return elevatorData
}