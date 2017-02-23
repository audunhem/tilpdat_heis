package driver

const MOTOR_SPEED = 2800

var ButtonLightChannels = [N_FLOORS][N_BUTTONS]int{
    {LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
    {LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
    {LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
    {LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var ButtonChannels = [N_FLOORS][N_BUTTONS]int{
    {BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
    {BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
    {BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
    {BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func InitElevator() {
    success := IoInit()
    if success == 0 {
        panic("Unable to initialize elevator hardware!")
    }
    for i := 0; i < N_FLOORS; i++ {
        for b := ButtonType(0); b < N_BUTTONS; b++ {
            SetButtonLamp(b, i, 0)
        }
    }

    SetStopLamp(0)
    SetDoorOpenLamp(0)
    SetFloorIndicator(0)
}

func SetMotorDirection(dir MotorDirection) {
    if dir == 0 {
        IoWriteAnalog(MOTOR, 0)
    } else if dir > 0 {
        IoClearBit(MOTORDIR)
        IoWriteAnalog(MOTOR, MOTOR_SPEED)
    } else if dir < 0 {
        IoSetBit(MOTORDIR)
        IoWriteAnalog(MOTOR, MOTOR_SPEED)
    }
}

func SetButtonLamp(button ButtonType, floor int, value int) {
    if floor < 0 {
        panic("Negative floor value")
    }
    if floor > N_FLOORS {
        panic("Floor value too high")
    }
    if button < 0 {
        panic("Negative button value")
    }
    if button > N_BUTTONS {
        panic("Button value too high")
    }

    if value == 1 {
        IoSetBit(ButtonLightChannels[floor][button])
    } else {
        IoClearBit(ButtonLightChannels[floor][button])
    }
}

func SetFloorIndicator(floor int) {
    if floor < 0 {
        panic("Negative floor value")
    }
    if floor > N_FLOORS {
        panic("Floor value too high")
    }

    // Binary encoding. One light must always be on.
    if floor&0x02 == 0x02 {
        IoSetBit(LIGHT_FLOOR_IND1)
    } else {
        IoClearBit(LIGHT_FLOOR_IND1)
    }

    if floor&0x01 == 0x01 {
        IoSetBit(LIGHT_FLOOR_IND2)
    } else {
        IoClearBit(LIGHT_FLOOR_IND2)
    }

}

func SetDoorOpenLamp(value int) {
    if value == 1 {
        IoSetBit(LIGHT_DOOR_OPEN)
    } else {
        IoClearBit(LIGHT_DOOR_OPEN)
    }
}

func SetStopLamp(value int) {
    if value == 1 {
        IoSetBit(LIGHT_STOP)
    } else {
        IoClearBit(LIGHT_STOP)
    }
}

func GetOrderButtonSignal(button ButtonType, floor int) int {
    if floor < 0 {
        panic("Negative floor value")
    }
    if floor > N_FLOORS {
        panic("Floor value too high")
    }
    if button < 0 {
        panic("Negative button value")
    }
    if button > N_BUTTONS {
        panic("Button value too high")
    }

    return IoReadBit((ButtonChannels[floor][button]))
}

func GetFloorSensorSignal() int {
    //must be changed if more floors
    if IoReadBit(SENSOR_FLOOR1) == 1 {
        return 0
    } else if IoReadBit(SENSOR_FLOOR2) == 1 {
        return 1
    } else if IoReadBit(SENSOR_FLOOR3) == 1 {
        return 2
    } else if IoReadBit(SENSOR_FLOOR4) == 1 {
        return 3
    } else {
        return -1
    }
}

func GetMotorDirection() MotorDirection {

    if IoReadAnalog(MOTOR) == 0 {
        return DirnStop
    } else if IoReadBit(MOTORDIR) == 1 {
        return DirnDown
    } else {
        return DirnUp
    }
}

func GetStopSignal() int {
    return IoReadBit(STOP)
}

func GetObstructionSignal() int {
    return IoReadBit(OBSTRUCTION)
}

func GetOpenDoor() int { //Sjekk om denne funker!!
    return IoReadBit(LIGHT_STOP)
}
