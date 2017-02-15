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
)

func init() {

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

  //Now lets assignetworkInit := make(chan peers.PeerUpdate)

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



func runNetwork(chan elevatorData updateTx, chan elevatorData updateRx, chan newOrder orderTx, chan newOrder orderRx)
  // First we need to asssign an ID to the elevator. We assume
  // That there can only be N_ELEV elevators at any time
  



  go bcast.Transmitter(16569, messageTx)
  go bcast.Receiver(16569, messageRx)

  go bcast.Transmitter(16568, orderTx)
  go bcast.Receiver(16568, orderRx)

  for {}


}
