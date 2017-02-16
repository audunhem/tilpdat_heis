package driver

const N_ELEV = 3

const N_FLOORS = 4

// Number of buttons (and corresponding lamps) on a per-floor basis
const N_BUTTONS = 3

type MotorDirection int8

const (
	DirnDown = -1 + iota
	DirnStop
	DirnUp
)

type ButtonType int

const (
	ButtonCallUp = iota
	ButtonCallDown
	ButtonCommand
)

type ElevatorStatus int

const (
	StatusIdle = iota
	StatusDoorOpen
	StatusMoving
)

type ElevatorData struct {
	Floor     int
	Direction MotorDirection
	Orders    [N_FLOORS][N_BUTTONS]int
	Status    ElevatorStatus
	Initiated int
}
