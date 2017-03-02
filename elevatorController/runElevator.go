package elevatorController

import (
	. "./../driver"
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
	initializedData.Initiated = true
	return initializedData
}

//Erling prøver å lage en ny versjon av denne

//TODÒ 23 feb: Skrive alt bedre, finne bedre måte for å
func ReadAllSensors2(arriveAtFloorCh chan int, externalButtonCh chan ElevatorOrder, internalButtonCh chan int) {

	currentFloor := GetFloorSensorSignal()
	//Lager variabel for å unngå å oppfatte tastetrykk flere ganger
	lastButtonPressed := -1

	for {
		//Vi ønsker kun beskjed hvis vi når en NY etasje! SKRIV DENNE PÅ EN BEDRE MÅTE, VI GJØR TRE KALL TIL GETFLOORSENSORSIGNAL
		if GetFloorSensorSignal() != currentFloor && GetFloorSensorSignal() >= 0 {
			currentFloor = GetFloorSensorSignal()
			arriveAtFloorCh <- currentFloor
		}

		//Looper gjennom alle EKSTERNE knapper
		for i := 0; i < N_FLOORS; i++ {
			for j := 0; j < 2; j++ {
				if GetOrderButtonSignal(ButtonType(j), i) == 1 {
					if lastButtonPressed != 2*i+j {
						lastButtonPressed = 2*i + j
						externalButtonCh <- ElevatorOrder{i, j, "-1"}
					}

				}
			}
		}

		//Looper gjennom alle INTERNE knapper

		for i := 0; i < N_FLOORS; i++ {
			if GetOrderButtonSignal(ButtonType(2), i) == 1 {
				if lastButtonPressed != N_FLOORS*2+i {
					lastButtonPressed = N_FLOORS*2 + i
					internalButtonCh <- i
					//Send info på internalButtonCh
				}
			}
		}

	}

	//Dette er egentlig alt denne funksjonen bør gjøre. Vi må finne på en god løsning på utfordringen av polling av knapper. Hvordan fungerer det egentlig?
	//Vil vi sende 1000 beskjeder om trykket inn knapp dersom en knapp holdes inn i 100ms?? MEst sannsynlig ikke

}

func getMacAddr() string {

	  var currentNetworkHardwareName string

	  interfaces, _ := net.Interfaces()
	  for _, interf := range interfaces {
	    currentNetworkHardwareName = interf.Name

	  }

	  // extract the hardware information base on the interface name
	  // capture above
	  netInterface, err := net.InterfaceByName(currentNetworkHardwareName)

	  if err != nil {
	    fmt.Println(err)
	  }

	  macAddress := netInterface.HardwareAddr

	  return macAddress.String()
}




/*
func ReadAllSensors(previousData ElevatorData, updatedDataFSM chan ElevatorData, currentFloorChannel chan int /*currentDirection chan MotorDirection,, newOrderButtonTypeChannel chan ButtonType, newOrderFloorChannel chan int) {
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
		if previousData.Initiated != true {
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

*/
