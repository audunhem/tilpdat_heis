package elevator

import (
	"log"
	"fmt"
 	"net"
	"bytes"
	"project_elevator/elevator/driver"
)

type ElevatorType int
const (
	ET_Comedi = iota
	ET_Simulation
)

var conn net.Conn
var elevatorType ElevatorType = ET_Simulation

func ElevatorInit(ip string, port int, et ElevatorType, elevatorInitiated int) {
	elevatorType = et

	switch elevatorType {
	case ET_Comedi:
		success := driver.Init()
		if success == 0 {
			panic("Hardware initiation failed")
		}

		for floor := 0; floor < N_FLOORS; floor ++ {
			for btn := Button(0); btn < N_BUTTONS; btn++ {
				ElevatorWriteOrderButtonLight(floor, btn, 0)
			}
		}

		ElevatorWriteStopButtonLight(0)
		ElevatorWriteDoorLight(0)
		ElevatorWriteFloorIndicator(0)

	case ET_Simulation:
		var buffer bytes.Buffer
		buffer.WriteString(ip)
		buffer.WriteString(":")
		buffer.WriteString(fmt.Sprintf("%d", port))
		newConn, err := net.Dial("tcp", buffer.String())

		if err != nil {
			panic("Unable to connect to elevator")
		}

		conn = newConn;

		if elevatorInitiated == 0 {
			msg := [4]byte{0}
			fmt.Fprintf(conn, string(msg[:]))
		}
	}
}

var floorSensorChannels = [N_FLOORS]int {
	SENSOR_FLOOR1,
	SENSOR_FLOOR2,
	SENSOR_FLOOR3,
	SENSOR_FLOOR4,
}

func ElevatorReadFloorSensor() int {
	switch elevatorType {
	case ET_Comedi:
		for f := 0; f < N_FLOORS; f++ {
			if driver.ReadBit(floorSensorChannels[f]) > 0 {
				return f
			}
		}
		return -1

	case ET_Simulation:
		msg := [4]byte{7}
		fmt.Fprintf(conn, string(msg[:]))

		buf := make([]byte, 4)
		_, err := conn.Read(buf)

		if err != nil {
			fmt.Println(err)
			return 0
		}

		if buf[1] == 1 { // Elevator is at some floor
			return int(buf[2]) // Floor
		} else {
			return -1
		}
	}
	return -2
}

var buttonChannels = [N_FLOORS][N_BUTTONS]int {
    {BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
    {BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
    {BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
    {BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func ElevatorReadOrderButton(floor int, button Button) int {
	switch elevatorType {
	case ET_Comedi:
		if floor < 0 || floor >= N_FLOORS {
			log.Println("Error: invalid floor in ElevatorReadRequestButton - " + string(floor))
			return 0
		}

		if button < 0 || button >= N_BUTTONS {
			log.Println("Error: invalid button in ElevatorReadRequestButton - " + string(button))
			return 0
		}

		return driver.ReadBit(buttonChannels[floor][button])

	case ET_Simulation:
		msg := [4]byte{6, byte(button), byte(floor)}
		fmt.Fprintf(conn, string(msg[:]))

		buf := make([]byte, 4)
		_, err := conn.Read(buf)

		if err != nil {
			log.Println(err)
			return 0
		}

		return int(buf[1])
	}
	return 0
}

func ElevatorReadStopButton() int {
	switch elevatorType {
	case ET_Comedi:
		return driver.ReadBit(STOP)
	
	case ET_Simulation:
		msg := [4]byte{8}
		fmt.Fprintf(conn, string(msg[:]))

		buf := make([]byte, 4)
		_, err := conn.Read(buf)

		if err != nil {
			log.Println(err)
			return 0
		}

		return int(buf[1])
	}
	return 0
}

func ElevatorReadObstruction() int {
	switch elevatorType {
	case ET_Comedi:
		return driver.ReadBit(OBSTRUCTION)

	case ET_Simulation:
		msg := [4]byte{9}
		fmt.Fprintf(conn, string(msg[:]))

		buf := make([]byte, 4)
		_, err := conn.Read(buf)

		if err != nil {
			log.Println(err)
			return 0
		}

		return int(buf[1])
	}
	return 0
}

func ElevatorWriteFloorIndicator(floor int) {
	switch elevatorType {
	case ET_Comedi:
		if floor < 0 || floor >= N_FLOORS {
			log.Println("Error: invalid floor in ElevatorWriteFloorIndicator")
		}

		if (floor & 0x02) != 0 {
			driver.SetBit(LIGHT_FLOOR_IND1)
		} else {
			driver.ClearBit(LIGHT_FLOOR_IND1)
		}

		if (floor & 0x01) != 0 {
			driver.SetBit(LIGHT_FLOOR_IND2)
		} else {
			driver.ClearBit(LIGHT_FLOOR_IND2)
		}

	case ET_Simulation:
		msg := [4]byte{3, byte(floor)}
		fmt.Fprintf(conn, string(msg[:]))
	}
}

var buttonLightChannels = [N_FLOORS][N_BUTTONS]int {
    {LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
    {LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
    {LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
    {LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

func ElevatorWriteOrderButtonLight(floor int, button Button, value int) {

	switch elevatorType {
	case ET_Comedi:
		if floor < 0 || floor >= N_FLOORS {
			fmt.Println("Error: invalid floor in ElevatorReadRequestButton - " + string(floor))
			return
		}

		if button < 0 || button >= N_BUTTONS {
			fmt.Println("Error: invalid button in ElevatorReadRequestButton - " + string(button))
			return
		}

		if value != 0 {
			driver.SetBit(buttonLightChannels[floor][button])
		} else {
			driver.ClearBit(buttonLightChannels[floor][button])
		}

	case ET_Simulation:
		msg := [4]byte{2, byte(button), byte(floor), byte(value)}
		fmt.Fprintf(conn, string(msg[:]))
	}
}

func ElevatorWriteDoorLight(value int) {
	switch elevatorType {
	case ET_Comedi:
		if value != 0 {
			driver.SetBit(LIGHT_DOOR_OPEN)
		} else {
			driver.ClearBit(LIGHT_DOOR_OPEN)
		}

	case ET_Simulation:
		msg := [4]byte{4, byte(value)}
		fmt.Fprintf(conn, string(msg[:]))
	}
}

func ElevatorWriteStopButtonLight(value int) {
	switch elevatorType {
	case ET_Comedi:
		if value != 0 {
			driver.SetBit(LIGHT_STOP)
		} else {
			driver.ClearBit(LIGHT_STOP)
		}

	case ET_Simulation:
		msg := [4]byte{5, byte(value)}
		fmt.Fprintf(conn, string(msg[:]))
	}
}

func ElevatorWriteMotorDirection(dirn Dirn) {
	switch elevatorType {
	case ET_Comedi:
		switch dirn {
		case D_Up:
			driver.ClearBit(MOTORDIR)
			driver.WriteAnalog(MOTOR, 2800)
		case D_Down:
			driver.SetBit(MOTORDIR)
			driver.WriteAnalog(MOTOR, 2800)
		case D_Stop:
			driver.WriteAnalog(MOTOR, 0)
		default:
			driver.WriteAnalog(MOTOR, 0)
		}

	case ET_Simulation:
		msg := [4]byte{1, byte(dirn)}
		fmt.Fprintf(conn, string(msg[:]))
	}
}