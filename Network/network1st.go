package network1st

import (
  "./network/bcast"
  "./network/localip"
  "./network/peers"
  "flag"
  "fmt"
  "os"
  "time"
  "strconv"
  "../driver/"
)



func RunNetwork(chan elevatorData updateTx, chan elevatorData updateRx, chan newOrder orderTx, chan newOrder orderRx, chan peers.PeerUpdate peerUpdateCh)
  // First we need to asssign an ID to the elevator. We assume
  // That there can only be N_ELEV elevators at any time

  //We will use functionality provided by the Network-Go module

  var id string
  var elevAlive int
  var i int

  elevAlive  = 0


  //Assign a unique ID to the elevator
  id = fmt.Sprintf("%s-%d", localIP, os.Getpid())

  //This is to send the ALIVE-signals

  peerTxEnable := make(chan bool)

  go peers.Transmotter(15647, id, peerTxEnable)
  go peers.Receiver(15647, peerUpdateCh)


//We initialize contact. Lets wait 5secs (or until all elevators
// are up and running).

  for i<5 {
    select {
    case p := <- peerUpdateCh:
      elevAlive = len(p.Peers)
      fmt.Printf("Elevator update:\n")
      fmt.Printf("  Elevators:    %q\n", p.Peers)
      fmt.Printf("  New:      %q\n", p.New)
      fmt.Printf("  Lost:     %q\n", p.Lost)
      if elevAlive == N_ELEV {
        break
      }

    default:
      i++
      time.Sleep(1*time.Second)
    }
  }


  //If this is the only elevator alive,
  if {elevAlive == 1} {

    panic()
  }

  go bcast.Transmitter(16569, messageTx)
  go bcast.Receiver(16569, messageRx)

  go bcast.Transmitter(16568, orderTx)
  go bcast.Receiver(16568, orderRx)

  for {}
}



/*


func init() {


  //THIS FUNCTION IS NOT GOING TO BE USED

  //Should return the ID of the elevator and
  //number of elevators connected



  localIP, err := localip.LocalIP()
  if err != nil {
    fmt.Println(err)
    localIP = "DISCONNECTED"
  }

  //We start by assigning a "random" ID to the elevator
  id := strconv.Itoa(os.Getpid())

  //enable is set to true to work with Anders network module
  //Which needs a par

  go peers.Transmitter(15647, id, enable)

  i := 0

  for i<4 {
    select {
    case peers := <-peerUpdateCh:
      //If all elevators are connected
      if len(peers.Peers) == N_ELEVATORS {
        i=10
      }
      time.Sleep(1*time.second)
    }

  }

  //Now lets assig

  peerUpdateCh := make(chan peers.PeerUpdate)

  go peers.Receiver(15647, networkInit)
  go peers.Transmitter(15647, id, enable)

  i := 0

  for i<3 {
    select {
    case peers := <-peerUpdateCh:
      i++
      //If all elevators are connected
      if len(peers.Peers) == N_ELEVATORS {
        i=10
      }
    case:
      Time.Sleep(1*Time.second)
    }

  }

  //Now lets assign the IDs we compare the process
  //IDs and the lowest is 1 and so on we will check how
  //many of the process IDs are "higher" than ours

  i = 0
  elevID := 1


  for i<len(peers.Peers) {
    if (strconv.Atoi(peers.Peers[i])>strconv.Atoi(id)) {
      elevID++
    }
  }

  return elevID
}n the IDs we compare the process
  //IDs and the lowest is 1 and so on we will check how
  //many of the process IDs are "higher" than ours

  i = 0
  elevID := 1


  for i<len(peers.Peers) {
    if (strconv.Atoi(peers.Peers[i])>strconv.Atoi(id)) {
      elevID++
    }
  }

  return elevID
}
