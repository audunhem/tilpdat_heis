package Events

import (
  "../Network/network/peers"
  . "./../driver"
)

const THIS_ELEVATOR = 1

var NETWORK_DOWN = false //hvordan er det bra å bruke denne variablen

var ExternalOrderLights = make([]int, 0) //skal den være her eller i intern-delen?

var Elevators = make([]ElevatorData, N_ELEVATORS)

func CalculateSingleElevatorCost(elevator ElevatorData, order ElevatorOrder) int {
  if int(elevator.Direction) == order.Direction {
    switch elevator.Direction {
    case DirnUp:
      if order.Floor > elevator.Floor {
        return order.Floor - elevator.Floor
      } else {
        return (elevator.Floor-1)*2 + (elevator.Floor - order.Floor)
      }
    case DirnDown:
      if order.Floor < elevator.Floor {
        return elevator.Floor - order.Floor
      } else {
        return (elevator.Floor-1)*2 + (order.Floor - elevator.Floor)
      }
    }
  } else {
    switch elevator.Direction {
    case DirnUp:
      return 2*N_FLOORS - elevator.Floor - order.Floor
    case DirnDown:
      return (elevator.Floor - 1) + (order.Floor - 1)
    }
  }
  return -1
}

func FindBestElevator(order ElevatorOrder) {
  var minCost = 1<<15 - 1
  var elevatorNumber = -1
  for i := 0; i < N_ELEVATORS; i++ {
    if Elevators[i].Initiated {
      var thisCost = CalculateSingleElevatorCost(Elevators[i], order)
      if thisCost < minCost {
        minCost = thisCost
        elevatorNumber = i
      }
    }
  }
  PlaceExternalOrder(elevatorNumber, order)
}

func PlaceOrder(elevatorData ElevatorData, order ElevatorOrder) ElevatorData {
  elevatorData.Orders[order.Floor][order.Direction] = 1
  return elevatorData
}

func PlaceExternalOrder(elevatorNumber int, order ElevatorOrder) {
  if elevatorNumber == THIS_ELEVATOR && !NETWORK_DOWN {
    Elevators[elevatorNumber].Orders[order.Floor][order.Direction] = 1
    //sort.Ints(Elevators[elevatorNumber].ExternalOrders.Floor) må kunne sortere knappene
  } else {
    //net.SendNewOrder()
  }
}

func SuccessfulPlacementConfirmation(elevatorNumber int, order ElevatorOrder) bool {
  if Elevators[elevatorNumber].Orders[order.Floor][order.Direction] == 1 {
    return true
  }
  return false
}

//må lage noe som merker at en heis har falt ut. -- lages i nettverk

//tror det er best om bare en av heisene omfordeler ordre

/*func RedestributeExternalOrders(lostElevator ElevatorData) {
  for i := 0; i < len(lostElevator.ExternalOrders); i++ {
    FindBestElevator(lostElevator.ExternalOrders[i])
  }
}*/

//må skrive om

func DenyNewExternalOrders() {
  NETWORK_DOWN = true
}

func AllowNewExternalOrders() {
  NETWORK_DOWN = false
}

func UpdateElevatorData(elevator ElevatorData) {
  for i := 0; i < N_ELEVATORS; i++ {
    if Elevators[i].ID == elevator.ID {
      Elevators[i] = elevator
    }
  }
}

//------------------------------------------------------------------------------
//Lagt til av Erling

//In case of updated list of connected elevators
func EventUpdatedPeers(updatedConnectionData peers.PeerUpdate) {

  //Either we have fewer connections or more connections. Either way
  //we want to update our

}
