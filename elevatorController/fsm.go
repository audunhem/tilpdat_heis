package elevatorController

import (
/*. "./../driver"
"fmt"
"time" */
)

func fsmArriveAtFloor(elevatorData ) {


  if OrderCheckIfShouldStop(elevatorStruct)
    fsmStopAtFloor()
    elevatorStruct = OrderCompleted(floor,direction. elevatorStruct)
    elevatorStruct = OrderSetNextDirection(elevatorStruct)
}


func fsmExternalButtonPressed() {




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
