package elevatorController

import (
	. "./../Events"
	. "./../driver"
	"fmt"
	"time"
)

//Funksjoner som skal legges til ORDER MODULEN
//------------------------------------------------------------------------------------------------------------

func FsmArriveAtFloor(elevatorStruct ElevatorData, floor int) ElevatorData {
	elevatorData := elevatorStruct
	elevatorData.Floor = floor
	PrintOrderList(elevatorData)
	if CheckIfShouldStop(elevatorData) == true {
		FsmStopAtFloor()
		elevatorData = OrderCompleted(elevatorData)
		PrintOrderList(elevatorData)
		//elevatorData = OrderCompleted(floor, elevatorData.o,elevatorData)
		elevatorData = OrderSetNextDirection(elevatorData) //sett elevator til IDLE
		//hvis det ikke er flere ordre
	}

	return elevatorData

}

func FsmExternalButtonPressed(elevatorStruct ElevatorData, newButtonPressed ElevatorOrder) ElevatorData {

	elevatorData := elevatorStruct
	elevatorData = PlaceExternalOrder2(elevatorData, newButtonPressed)

	if elevatorData.Status == StatusIdle {
		fmt.Println("IDLE")
		elevatorData = OrderSetNextDirection(elevatorData)
	}

	return elevatorData

}

func FsmStopAtFloor() {
	SetMotorDirection(DirnStop)
	SetDoorOpenLamp(1)
	time.Sleep(500 * time.Millisecond)
	SetDoorOpenLamp(0)
}

func PrintOrderList(elevatorStruct ElevatorData) {
	for i := 0; i < N_FLOORS; i++ {
		for j := 0; j < N_BUTTONS; j++ {
			fmt.Printf("%d", elevatorStruct.Orders[i][j])
		}
		fmt.Printf("\n")

	}

	fmt.Printf("--------------------------------------------")
	fmt.Printf("\n")

}

func FsmInternalButtonPressed(elevatorStruct ElevatorData, floor int) ElevatorData {

	elevatorData := PlaceInternalOrder(elevatorStruct, floor)

	if elevatorData.Status == StatusIdle {

		elevatorData = OrderSetNextDirection(elevatorData)
	}

	return elevatorData

}

/*

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
