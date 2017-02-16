package elevatorController

import (
	. "./../driver"
	"fmt"
	"time"
)

func InitializeElevator() ElevatorData {
	//sett heisen i en etasje
	//oppdater structen.
	//sett initialisert til true
	InitElevator()

	if GetFloorSensorSignal() == -1 {
		SetMotorDirection(DirnUp)
		for GetFloorSensorSignal() == -1 {
		}
		SetMotorDirection(DirnStop)
	}
	var initializedData ElevatorData

	initializedData.Floor = GetFloorSensorSignal()
	initializedData.Direction = GetMotorDirection()
	initializedData.Status = 0
	initializedData.Initiated = 1
	return initializedData
}

func ReadAllSensors(updatedDataFSM chan ElevatorData, currentFloorChannel chan int /*currentDirection chan MotorDirection,*/, newOrderButtonTypeChannel chan ButtonType, newOrderFloorChannel chan int) {
	//check all sensors.
	//update data
	//set all lights
	var previousData ElevatorData
	var updatedData ElevatorData
	var currentFloor int

	var i int
	previousData = InitializeElevator()
	for {
		if previousData.Initiated != 1 {
			panic("ElevatorNotInitialized")
		}
		currentFloor = GetFloorSensorSignal()

		updatedData.Floor = currentFloor
		fmt.Println(updatedData.Floor)
		updatedData.Direction = GetMotorDirection()

		if GetMotorDirection() != 0 {
			updatedData.Status = 2
		} else if GetOpenDoor() == 1 {
			updatedData.Status = 1
			SetDoorOpenLamp(1)
		} else {
			updatedData.Status = 0
		}

		SetFloorIndicator(updatedData.Floor)
		i = i + 1
		if i == 10 {
			previousData = updatedData
			i = 0
		}
		updatedDataFSM <- updatedData
		currentFloorChannel <- currentFloor

		GetNewOrders(updatedData, updatedData, newOrderButtonTypeChannel, newOrderFloorChannel)
	}
}

func GetNewOrders(updatedData ElevatorData, previousData ElevatorData, newOrderButtonTypeChannel chan ButtonType, newOrderFloorChannel chan int) bool {
	var newOrderButtonType ButtonType
	var newOrderFloor int

	for floor := 0; floor < N_FLOORS; floor++ {
		for btn := ButtonType(0); btn < N_BUTTONS; btn++ {
			updatedData.Orders[floor][btn] = GetOrderButtonSignal(btn, floor)
			if previousData.Orders[floor][btn] != updatedData.Orders[floor][btn] {
				newOrderButtonType = btn
				newOrderFloor = floor
				previousData = updatedData
				return true
			}
		}
		newOrderButtonTypeChannel <- newOrderButtonType
		newOrderFloorChannel <- newOrderFloor
	}
	return false
}

func GoToFloor(floor int) bool {

	if (floor - GetFloorSensorSignal()) > 0 {
		SetMotorDirection(DirnUp)
	} else if (floor - GetFloorSensorSignal()) < 0 {
		SetMotorDirection(DirnDown)
	}
	for GetFloorSensorSignal() != floor {
	}
	SetMotorDirection(DirnStop)
	OpenDoors()

	return true
}

func OpenDoors() {
	if GetMotorDirection() != DirnStop {
		fmt.Println("Heisen har ikke stoppet")
	}
	SetDoorOpenLamp(1)
	time.Sleep(3 * time.Second)
	SetDoorOpenLamp(0)
}
