package Events

import (
  //"sort"
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
  ExternalOrders []Button //trenger noe for å vise retningen til ordren
  Direction int
  CurrentFloor int
  ID int
  Alive bool
}

var ExternalButtonLights = make([]int, 0) //skal den være her eller i intern-delen?

var Elevators = make([]Elevator, N_ELEVATORS)

func CalculateSingleElevatorCost(elevator Elevator, button Button) int{
  if elevator.Direction == button.direction {
      switch elevator.Direction {
      case DirnUp:
        if button.floor > elevator.CurrentFloor{
          return button.floor - elevator.CurrentFloor
        } else {
          return (elevator.CurrentFloor-1)*2 + (elevator.CurrentFloor-button.floor)
        }
      case DirnDown:
      if button.floor < elevator.CurrentFloor{
        return elevator.CurrentFloor - button.floor
      } else {
        return (elevator.CurrentFloor-1)*2 + (button.floor - elevator.CurrentFloor)
      }
    }
    } else {
      switch elevator.Direction {
      case DirnUp:
        return 2*N_FLOORS - elevator.CurrentFloor - button.floor
      case DirnDown:
        return (elevator.CurrentFloor-1) + (button.floor-1)
      }
    }
    return -1
}

func FindBestElevator(button Button){
  var minCost = 1<<15 - 1
  var elevatorNumber = -1
  for i := 0; i < N_ELEVATORS; i++ {
    if Elevators[i].Alive {
      var thisCost = CalculateSingleElevatorCost(Elevators[i], button)
      if  thisCost < minCost{
        minCost = thisCost
        elevatorNumber = i
      }
    }
  }
  PlaceExternalOrder(elevatorNumber, button)
}

func PlaceExternalOrder(elevatorNumber int, button Button){
  if elevatorNumber == THIS_ELEVATOR && !NETWORK_DOWN  {
    Elevators[elevatorNumber].ExternalOrders = append(Elevators[elevatorNumber].ExternalOrders,button)
    //sort.Ints(Elevators[elevatorNumber].ExternalOrders.floor) må kunne sortere knappene
  } else {
    //net.SendNewOrder()
  }
}

func SuccessfulPlacementConfirmation(elevatorNumber int, button Button) bool{
  for i := 0; i < len(Elevators[elevatorNumber].ExternalOrders); i++{
      if Elevators[elevatorNumber].ExternalOrders[i].floor == button.floor{
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

//gitgitgit


