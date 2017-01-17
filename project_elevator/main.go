package main

import(
	"time"
	"log"
	"net"
	"project_elevator/processController"
)

func main() {
	
	// Initiating UDP address for backup process
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:20000")
	if err != nil {log.Fatal(err)}

	// Backup listening to primary
	udpListen, err := net.ListenUDP("udp", udpAddr)
	if err != nil {log.Fatal(err)}
	
	// Last elevatorData retreived from primary
	elevatorBackup := processController.StartBackupProcess(udpListen)
	
	// Backup taking over as primary. Closing listening connection.
	udpListen.Close()
	
	// Initiating UDP address for primary process.
	udpAddr, err = net.ResolveUDPAddr("udp","localhost:20000")
	if err != nil {log.Fatal(err)}

	// Setting up connection to backup.
	udpBroadcast, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {log.Fatal(err)}
	
	// Starting primary process
	time.Sleep(100*time.Millisecond)
	processController.StartPrimaryProcess(elevatorBackup, udpBroadcast)
	
	// Closing UDP connection when primary process fails.
	udpBroadcast.Close()
}