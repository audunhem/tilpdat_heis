package Events

import (
  //"sort"
  "../Network/network/peers"
)

const THIS_ELEVATOR = 1
const N_ELEVATORS = 3
const N_FLOORS = 4

var NETWORK_DOWN = false //hvordan er det bra å bruke denne variablen



type Order struct{
  ID string
  floor int
  direction int
}

type Elevator struct{
  InternalOrders []int
  ExternalOrders []Order //trenger noe for å vise retningen til ordren
  Direction int
  CurrentFloor int
  ID int
  Alive bool
}

var ExternalOrderLights = make([]int, 0) //skal den være her eller i intern-delen?

var Elevators = make([]Elevator, N_ELEVATORS)

func CalculateSingleElevatorCost(elevator Elevator, order Order) int{
  if elevator.Direction == Order.direction {
      switch elevator.Direction {
      case DirnUp:
        if order.floor > elevator.CurrentFloor{
          return order.floor - elevator.CurrentFloor
        } else {
          return (elevator.CurrentFloor-1)*2 + (elevator.CurrentFloor-order.floor)
        }
      case DirnDown:
      if order.floor < elevator.CurrentFloor{
        return elevator.CurrentFloor - order.floor
      } else {
        return (elevator.CurrentFloor-1)*2 + (order.floor - elevator.CurrentFloor)
      }
    }
    } else {
      switch elevator.Direction {
      case DirnUp:
        return 2*N_FLOORS - elevator.CurrentFloor - order.floor
      case DirnDown:
        return (elevator.CurrentFloor-1) + (order.floor-1)
      }
    }
    return -1
}

func FindBestElevator(order Order){
  var minCost = 1<<15 - 1
  var elevatorNumber = -1
  for i := 0; i < N_ELEVATORS; i++ {
    if Elevators[i].Alive {
      var thisCost = CalculateSingleElevatorCost(Elevators[i], order)
      if  thisCost < minCost{
        minCost = thisCost
        elevatorNumber = i
      }
    }
  }
  PlaceExternalOrder(elevatorNumber, order)
}

func PlaceExternalOrder(elevatorNumber int, order Order){
  if elevatorNumber == THIS_ELEVATOR && !NETWORK_DOWN  {
    Elevators[elevatorNumber].ExternalOrders = append(Elevators[elevatorNumber].ExternalOrders,order)
    //sort.Ints(Elevators[elevatorNumber].ExternalOrders.floor) må kunne sortere knappene
  } else {
    //net.SendNewOrder()
  }
}

func SuccessfulPlacementConfirmation(elevatorNumber int, order Order) bool{
  for i := 0; i < len(Elevators[elevatorNumber].ExternalOrders); i++{
      if Elevators[elevatorNumber].ExternalOrders[i].floor == order.floor{
        return true
      }
  }
  return false
}

//må lage noe som merker at en heis har falt ut. -- lages i nettverk

//tror det er best om bare en av heisene omfordeler ordre

func RedestributeExternalOrders(lostElevator Elevator){
  for i := 0; i < len(lostElevator.ExternalOrders); i++ {
    FindBestElevator(lostElevator.ExternalOrders[i])
  }
}

func DenyNewExternalOrders(){
  NETWORK_DOWN = true
}

func AllowNewExternalOrders(){
  NETWORK_DOWN = false
}

func UpdateElevatorData(elevator Elevator){
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

func main(){
  return 0
}
