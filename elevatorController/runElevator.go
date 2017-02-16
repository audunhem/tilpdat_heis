package elevatorController


func InitilizeElevator() ElevatorData{
	//sett heisen i en etasje
	//oppdater structen. 
	//sett initialisert til true
	InitElevator()

	if (GetFloorSensorSignal() == -1) {
		SetMotorDirection(DirnUp)
		for (GetFloorSensorSignal() == -1) {}
		SetMotorDirection(DirnStop)
	}
	var initializedData ElevatorData

	initializedData.Floor=GetFloorSensorSignal()
	initializedData.Direction=GetMotorDirection()
	initializedData.Status=0
	initializedData.Initiated=1
	return initializedData
}

}
func ReadAllSensors(updatedDataFSM chan ElevatorData, currentFloor chan ElevatorData, currentDirection chan MotorDirection,newOrderButtonType chan ButtonType, newOrderFloor chan int){
	//check all sensors. 
	//update data
	//set all lights
	previousData=InitializeElevator()
	for{

		if previousData.Initiated!=1
			panic("ElevatorNotInitialized")

		updatedDataFSM.Floor=GetFloorSensorSignal()
		updatedDataFSM.Direction=GetFloorSensorSignal()

		if GetMotorDirection()!=0{
			updatedDataFSM.Status=2
		}else if GetOpenDoor()==1{
			updatedDataFSM.Status=1
			SetOpenDoorLamp(1)
		}
		else{updatedDataFSM.Status=0}
		
		SetFloorIndicator(updatedDataFSM.Floor)

		GetNewOrders(updatedDataFSM, previousData, newOrderButtonType, newOrderFloor) 
	}
}

func GetNewOrders(updatedDataFSM chan ElevatorData, previousData ElevatorData, newOrderButtonType chan ButtonType, newOrderFloor chan int) bool {
	for floor := 0; floor < N_FLOORS; floor++{
		for btn := elevator.ButtonType(0); btn < N_BUTTONS; btn++{
			elevator.GetOrderButtonSignal(floor, btn, currentData.Orders[floor][btn])
			if previousData.Orders[floor][btn]=!updatedDataFSM.Orders[floor][btn]{
				newOrderButtonType=btn
				newOrderFloor=floor
				previousData=updatedDataFSM
				return true
		}
	}
	return false
}}



