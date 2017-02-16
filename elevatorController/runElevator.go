package elevatorController

import (
	. "./../driver"
	"fmt"
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

//Erling prøver å lage en ny versjon av denne
func ReadAllSensors2(arriveAtFloorCh chan int, externalButtonCh chan int, internalButtonCh chan int) {

	currentFloor := GetFloorSensorSignal()

	for {
		//Vi ønsker kun beskjed hvis vi når en NY etasje!
		if GetFloorSensorSignal() != currentFloor {
			currentFloor = GetFloorSensorSignal()
			//Her polles floor sensoren to ganger, er det unødvendig?
			//send info på channel
		}

		//Looper gjennom alle EKSTERNE knapper
		for i := 0; i < N_FLOORS; i++ {
			for j := 0; j < 1; j++ {
				if GetOrderButtonSignal(ButtonChannels[i][j]) == 1 {
					//Send info på externalButtonCh
				}
			}
		}

		//Looper gjennom alle INTERNE knapper

		for i := 0; i < N_FLOORS; i++ {
			if GetOrderButtonSignal(ButtonChannels[i][ButtonCommand]) == 1 {
				//Send info på internalButtonCh
			}
		}

	}

	//Dette er egentlig alt denne funksjonen bør gjøre. Vi må finne på en god løsning på utfordringen av polling av knapper. Hvordan fungerer det egentlig?
	//Vil vi sende 1000 beskjeder om trykket inn knapp dersom en knapp holdes inn i 100ms?? MEst sannsynlig ikke

}

func ReadAllSensors(previousData ElevatorData, updatedDataFSM chan ElevatorData, currentFloorChannel chan int /*currentDirection chan MotorDirection,*/, newOrderButtonTypeChannel chan ButtonType, newOrderFloorChannel chan int) {
	//check all sensors.
	//update data
	//set all lights
	fmt.Println("Begin reading sensors")
	//var previousData ElevatorData
	var currentFloor int

	var updatedData ElevatorData

	var i int
	//previousData = InitializeElevator()
	for {
		if previousData.Initiated != 1 {
			panic("ElevatorNotInitialized")
		}
		currentFloor = GetFloorSensorSignal()

		updatedData.Floor = currentFloor
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
		fmt.Println("okok")
		updatedDataFSM <- updatedData
		currentFloorChannel <- currentFloor
		fmt.Println("etter")
		//GetNewOrders(updatedData, updatedData, newOrderButtonTypeChannel, newOrderFloorChannel)
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
