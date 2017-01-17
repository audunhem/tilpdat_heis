package network

import (
	"net"
	"log"
)

const (
	srvAddr         = "239.100.100.100:30001"
	maxDatagramSize = 8192
)

func Init(outgoingMessage chan []byte, receivedMessages chan []byte, killChannel chan int){
	go MulticastMessage(srvAddr, outgoingMessage, killChannel)
	go ReceiveMessage(srvAddr, receivedMessages)
}

func MulticastMessage(networkAddr string, outgoingMessage chan []byte, killChannel chan int){
	addr, err := net.ResolveUDPAddr("udp", networkAddr)
	if err != nil {
		log.Println(err)

		select {
		case killChannel <- 1:
		default:
			log.Println("Killchannel in network is full")	
		}

		return
	}

	connection, err := net.DialUDP("udp", nil, addr)
	if err != nil{
		log.Println(err)

		select{
		case killChannel <- 1:
		default:
			log.Println("Killchannel in network is full")	
		}

		return
	}

	for {
		select{
		case messageOut := <- outgoingMessage:
			_, err := connection.Write(messageOut)
			if err != nil {
				log.Println(err, "Elevator will now run in offline mode.")

				select{
				case killChannel <- 1:
				default:
					log.Println("Killchannel in network is full")	
				}

				return
			}
		}
	}
}

func ReceiveMessage(networkAddr string, receivedMessages chan []byte){
	addr, err := net.ResolveUDPAddr("udp", networkAddr)
	if err != nil{
		log.Println(err)
		return
	}

	listeningConnection, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil{
		log.Println(err)
		return
	}

	listeningConnection.SetReadBuffer(maxDatagramSize)
	for {
		messageBuffer := make([]byte, maxDatagramSize)
		n, _, err := listeningConnection.ReadFromUDP(messageBuffer)

		if err != nil{
			log.Println(err)
		}
		
		select{
		case receivedMessages <- messageBuffer[:n]:
		default:
			log.Println("receivedMessages channel is full")
		}
	}
}