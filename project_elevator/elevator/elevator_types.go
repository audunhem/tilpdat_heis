package elevator

const N_FLOORS = 9
const N_BUTTONS = 3

type Dirn int8 
const ( 
	D_Down Dirn = -1 + iota
	D_Stop
	D_Up
)

type Button byte
const (
	B_HallUp = iota
	B_HallDown
	B_Cab
)

type ElevatorBehaviour int
const (
	EB_Idle = iota
	EB_DoorOpen
	EB_Moving
)

type ElevatorData struct {
	Floor 		int
	Direction	Dirn
	Orders	 	[N_FLOORS][N_BUTTONS]int
	Behaviour 	ElevatorBehaviour
	Initiated 	int
}