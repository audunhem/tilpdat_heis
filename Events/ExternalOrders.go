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
  if int(elevator.Direction) == order.Direction { //her blir det krøll
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
  var elevatorNumber = -1 //kanksje fint å bruke ID her?
  for i := 0; i < N_ELEVATORS; i++ {
    if Elevators[i].Initiated {
      var thisCost = CalculateSingleElevatorCost(Elevators[i], order)
      if thisCost < minCost {
        minCost = thisCost
        order.ElevatorID = Elevators[i].ID
      }
    }
  }
  PlaceExternalOrder(order) //kan bare bruke ID-en til ordren
}

//Dette må også ordnes 23feb
func PlaceExternalOrder2(elevatorData ElevatorData, order ElevatorOrder) ElevatorData {
  elevatorData.Orders[order.Floor][order.Direction] = 1
  return elevatorData
}

func PlaceInternalOrder(elevatorData ElevatorData, floor int) ElevatorData {

  elevatorData.Orders[floor][ButtonType(2)] = 1

  return elevatorData
}

func PlaceExternalOrder(order ElevatorOrder) {
  if order.ElevatorID != nil {
    for i := 0 ; i < N_ELEVATORS; i++ {
      if order.ElevatorID == Elevators[i].ID {
        Elevators[i].Orders[order.Floor][order.Direction] = 1
      }
    }
  } else {
    FindBestElevator(order)
  }
}

//trenger noe lurt som sikrer at eksterne lys blir riktig

func SuccessfulPlacementConfirmation(elevatorNumber int, order ElevatorOrder) bool {
  if Elevators[elevatorNumber].Orders[order.Floor][order.Direction] == 1 {
    return true
  }
  return false
}

//må lage noe som merker at en heis har falt ut. -- lages i nettverk

//tror det er best om bare en av heisene omfordeler ordre

func RedestributeExternalOrders(lostElevator ElevatorData) {
  for i := 0; i < N_FLOORS; i++ {
    for j := 0; j < 2 ; j++ {
      if lostElevator.Orders[i][j] == 1 {
        newOrder := Order{i,j,nil}
        FindBestElevator(newOrder)
      }
    }
  }
}

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
