package elevatorController

import (
	. "./../driver"
	/*"fmt"
	  "time" */)

//Funksjoner som skal legges til ORDER MODULEN
//------------------------------------------------------------------------------------------------------------

func OrderSetNextDirection(elevatorStruct ElevatorData) ElevatorData {
	elevatorData := elevatorStruct
	check := 0

	if elevatorData.Status == StatusIdle {
		for i := 0; i < N_FLOORS; i++ {
			for j := 0; j < N_BUTTONS; j++ {
				if elevatorData.Orders[i][j] == 1 {
					if elevatorData.Floor < i {
						elevatorData.Direction = DirnUp
						SetMotorDirection(DirnUp)
						elevatorData.Status = StatusMoving
					} else if elevatorData.Floor > i {
						elevatorData.Direction = DirnDown
						SetMotorDirection(DirnDown)
						elevatorData.Status = StatusMoving
					} else if elevatorData.Floor == i {
						elevatorData = fsmArriveAtFloor(i, elevatorData)
					}
				}

			}
		}

	} else if elevatorData.Direction == DirnUp {
		for i := elevatorData.Floor; i < N_FLOORS; i++ {
			for j := 0; j < N_BUTTONS; j++ {
				if elevatorData.Orders[i][j] == 1 {
					SetMotorDirection(DirnUp)
					check = 1
				}
			}
		}

		if check == 0 {
			for i := 0; i < elevatorData.Floor; i++ {
				for j := 0; j < N_BUTTONS; j++ {
					if elevatorData.Orders[i][j] == 1 {
						SetMotorDirection(DirnDown)
						elevatorData.Direction = DirnDown
					}
				}
			}
		} else if elevatorData.Direction == DirnDown {

			for i := 0; i < elevatorData.Floor; i++ {
				for j := 0; j < N_BUTTONS; j++ {
					if elevatorData.Orders[i][j] == 1 {
						SetMotorDirection(DirnDown)
						check = 1
					}
				}
			}

			if check == 0 {
				for i := elevatorData.Floor; i < N_FLOORS; i++ {
					for j := 0; j < N_BUTTONS; j++ {
						if elevatorData.Orders[i][j] == 1 {
							SetMotorDirection(DirnUp)
							elevatorData.Direction = DirnUp
						}
					}
				}
			}
		}
	}

	return elevatorData
}

func fsmArriveAtFloor(floor int, elevatorStruct ElevatorData) ElevatorData {
	elevatorData := elevatorStruct

	/*
		if OrderCheckIfShouldStop(elevatorStruct) == 1 {
			//fsmStopAtFloor()
			//elevatorData = OrderCompleted(floor, elevatorData.o,elevatorData)
			elevatorData = OrderSetNextDirection(elevatorData) //sett elevator til IDLE
			//hvis det ikke er flere ordre
		}
	*/

	return elevatorData

}

/*
func fsmExternalButtonPressed(elevatorStruct ElevatorData, newButtonPressed elevatorOrder) ElevatorData {

	if true {
		elevatorData = elevatorStruct
		elevatorData = OrderAddOrder(newButtonPressed, elevatorStruct)

	}
	if !elevatorData.ElevatorStatus {
		elevatorData = OrderSetNextDirection(elevatorData)
	}

	return elevatorData

}

func fsmInternalButtonPressed(elevatorStruct ElevatorData, newButtonPressed ElevatorOrder) ElevatorData {

	elevatorData := OrderAddOrder(newButtonPressed, elevatorStruct)

	if !elevatorData.ElevatorStatus {

		elevatorData = OrderSetNextDirection(elevatorData)
	}

	return elevatorData

}

func goDown() {}

func goUp() {}

func openDoors() {}

func stop() {}

func readAllSensors() {}

//DETTE ER VAR FØRSTE UTKASTET PRØVER PÅ NYTT
/*


func GoToFloor(floor int, updatedData *ElevatorData) bool {
	fmt.Println("GoToFloor")
	if (floor - (*updatedData).Floor ) > 0 {
		SetMotorDirection(DirnUp)
		fmt.Println("MotorDirectionUp")
	} else if (floor - (*updatedData).Floor ) < 0 {
		SetMotorDirection(DirnDown)
		fmt.Println("MotorDirectionDown")
	}
	for (*updatedData).Floor != floor {
	}
	fmt.Println("GoToFloorEnd")
	SetMotorDirection(DirnStop)
	OpenDoors()
	fmt.Println("GoToFloorEnd")

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

func GoUp...() {

}


/*
import (
	."./../driver"
	"fmt"
	"time"

)

func GoToFloor(floor int){
	//Takes the elevator to _floor_, opens the door for 3 secs and closes it. Returns 1 on success



	if (GetFloorSensorSignal() == -1) {
		SetMotorDirection(DirnUp)

			for (GetFloorSensorSignal() == -1) {}
			SetMotorDirection(DirnStop)
	}




	if (floor - GetFloorSensorSignal()) > 0 {
		SetMotorDirection(DirnUp)
	} else if (floor - GetFloorSensorSignal()) < 0 {
		SetMotorDirection(DirnDown)
	}

	fmt.Println(GetMotorDirection())
	for (GetFloorSensorSignal() != floor) {
	}




	SetMotorDirection(DirnStop)
	OpenDoors()



*/

//else if (floor == GetFloorSensorSignal())

/*
}


func OpenDoors() {
	if GetMotorDirection() != DirnStop {
		fmt.Println("gjør noe")
	}
	SetDoorOpenLamp(1)
	time.Sleep(3*time.Second)
	SetDoorOpenLamp(0)
}*/
