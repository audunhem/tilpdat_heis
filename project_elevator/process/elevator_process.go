package process

import(
		"time"
		"log"
		"net"
		"os/exec"
		"project_elevator/elevator"
		"project_elevator/fsm"
		"project_elevator/timer"
		"project_elevator/network"
		"project_elevator/messageHandler"
		"project_elevator/fileManager"
		"project_elevator/typeConverter"
)

var elevatorData elevator.ElevatorData

const N_FLOORS = elevator.N_FLOORS
const N_BUTTONS = elevator.N_BUTTONS

func StartPrimaryProcess(initialData elevator.ElevatorData, udpBroadcast *net.UDPConn){

	// Setting up new backup process in separate terminal
	newBackup := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run *.go")
	err := newBackup.Run()
	if err != nil {log.Fatal(err)}

	// Communication will start here
	outgoingMessages := make(chan []byte, 1024)
	incomingMessages := make(chan []byte, 1024) 
	killChannel := make(chan int, 1)
	buttonPressed := make(chan []byte, 1024)
	orderComplete := make(chan []byte, 1024)
	updateElevatorData := make(chan []byte, 1024)
	updateOrders := make(chan []byte, 1024)

	network.Init(outgoingMessages, incomingMessages, killChannel)
	messageHandler.Init(outgoingMessages, incomingMessages, killChannel, buttonPressed, orderComplete, updateElevatorData, updateOrders)

	RunElevator(initialData, udpBroadcast, buttonPressed, orderComplete, updateElevatorData, updateOrders)

}

func StartBackupProcess(udpListen *net.UDPConn) elevator.ElevatorData{
	listenChan := make(chan elevator.ElevatorData, 1); 
	var backupData elevator.ElevatorData

	// Start Listen as separate goroutine
	go Listen(listenChan, udpListen)

	// Select waits for one case, then sleeps. Listen then runs.
	for {
		select {
			case backupData = <- listenChan:
				time.Sleep(50*time.Millisecond)
			case <-time.After(1*time.Second):
				if backupData.Initiated != 0 {
					log.Println("Primary elevator failed. Backup taking over.")
				}
				return backupData
		}
	}
}

func Listen(listenChan chan elevator.ElevatorData, udpListen *net.UDPConn) {

	buffer := make([]byte, 1024)

	for {
		udpListen.ReadFromUDP(buffer[:])
		//if err != nil {log.Fatal(err)} 
		
		elevatorData, err := typeConverter.ConvertMessageToElevatorData(buffer)
		if err != nil {
			panic(err)
		} else {
			listenChan <- elevatorData
		}
		
		time.Sleep(100*time.Millisecond)
	}
	
}

func RunElevator(initialData elevator.ElevatorData, udpBroadcast *net.UDPConn, buttonPressed chan []byte, orderComplete chan []byte, updateElevatorData chan []byte, 
	updateOrders chan []byte) {

	elevatorData = initialData

	if elevatorData.Initiated != 1 { // First startup.
		orders, err := fileManager.ReadOrdersFromFile()
		if err != nil {
			log.Println(err)
			log.Println("No order data will be loaded.")
		}	else {
			elevatorData.Orders = orders
			log.Println("Stored orders were loaded from file.")
		}
	}

	elevator.ElevatorInit("localhost", 15657, elevator.ET_Simulation, elevatorData.Initiated) // Error handling?

	if elevator.ElevatorReadFloorSensor() == -1 {
		elevatorData = fsm.FsmOnInitBetweenFloors(elevatorData)
		elevatorData.Initiated = 1	// Necessary for using the simulator.
	}

	var prevO [N_FLOORS][N_BUTTONS] int 
	var prevF int

	for {
		{ // Check buttons
			
				for f := 0; f < N_FLOORS; f++ {
					for b := elevator.Button(0); b < N_BUTTONS; b++ {
						v := elevator.ElevatorReadOrderButton(f, b)
						if v != 0 && v != prevO[f][b] {
							elevatorData = fsm.FsmOnOrderButtonPressed(f, b, elevatorData, buttonPressed, orderComplete)
							select {
							case updateElevatorData <- typeConverter.ConvertElevatorDataToMessage(elevatorData):
							default:
								log.Println("Update elevator data channel full")
							}
							
						}
						prevO[f][b] = v
					}
				}
		}

		{ // Check floor sensors
			f := elevator.ElevatorReadFloorSensor()

			if f != -1 && f != prevF {
				elevatorData = fsm.FsmOnFloorArrival(f, elevatorData, orderComplete)
				select {
				case updateElevatorData <- typeConverter.ConvertElevatorDataToMessage(elevatorData):
				default:
					log.Println("Update elevator data channel full")
				}
			}
			prevF = f
		}

		{ // Timer
			if timer.TimerTimeout() {
				elevatorData = fsm.FsmOnDoorTimeout(elevatorData)
				select {
				case updateElevatorData <- typeConverter.ConvertElevatorDataToMessage(elevatorData):
				default:
					log.Println("Update elevator data channel full")
				}
				timer.TimerStop()
			}
		}

		{ // Update backup process
			msg := typeConverter.ConvertElevatorDataToMessage(elevatorData)
			udpBroadcast.Write(msg)
		}

		{ // Update ElevatorData

			select {
			case order := <- updateOrders:
				floor := int(order[0])
				buttonType := int(order[1])
				setValue := int(order[2])

				elevator.ElevatorWriteOrderButtonLight(floor, elevator.Button(buttonType), setValue)

				if elevatorData.Orders[floor][buttonType] != setValue {
					elevatorData.Orders[floor][buttonType] = setValue
					fileManager.WriteOrdersToFile(elevatorData)
					elevatorData = fsm.FsmOnOrdersUpdate(elevatorData, orderComplete)
				}

			default:
			}
		}
		time.Sleep(100*time.Millisecond)
	}

}

