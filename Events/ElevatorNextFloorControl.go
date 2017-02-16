package event

import (
  "./../../elevatorController"
)

const (
	DirnDown = -1 + iota
	DirnStop
	DirnUp
)

var thisElevator = Elevators[0]
//trenger noe ala thisElevator

func RecieveNewState(a,b,c,d chan int){
  select{
  case c <- x:
    //sett noe til noe
  case c2 <- y:
    //sett noe annet til noe annet
  }
}

func SendElevatorToNextFloor(newInternalOrder button){ //må vente til det kommer en intern ordre før man sender neste etasje
  switch{
    case newInternalOrder.floor > thisElevator.currentFloor: //directrion is up
      nextFloor = N_FLOORS
      for i := 0; i < thisElevator.InternalOrders; i++ {
        if (thisElevator.InternalOrders[i] > currentFloor) && (thisElevator.InternalOrders[i] < nextFloor) {
          nextFloor = thisElevator.InternalOrders[i]
        }
      }
      for j := 0; j < thisElevator.ExternalOrders; j++ {
        if (thisElevator.ExternalOrders[i] > currentFloor) && (thisElevator.ExternalOrders[i] < nextFloor) {
          if (thisElevator.ExternalOrders[i].direction == 1) || (thisElevator.ExternalOrders[i].floor == N_FLOORS){
            nextFloor = thisElevator.ExternalOrders[i]
          }
        }
      }
    case newInternalOrder.floor < thisElevator.currentFloor: //directrion is down
      nextFloor = 1
      for i := 0; i < thisElevator.InternalOrders; i++ {
        if (thisElevator.InternalOrders[i] > currentFloor) && (thisElevator.InternalOrders[i] > nextFloor) {
          nextFloor = thisElevator.InternalOrders[i]
        }
      }
      for j := 0; j < thisElevator.ExternalOrders; j++ {
        if (thisElevator.ExternalOrders[i] < currentFloor) && (thisElevator.ExternalOrders[i] > nextFloor) {
          if (thisElevator.ExternalOrders[i].direction == -1) || (thisElevator.ExternalOrders[i].floor == 1){
            nextFloor = thisElevator.ExternalOrders[i]
          }
        }
      }
  }
  GoToFloor(nextFloor)
}

func MotorOutOfOrder(){
  //trenger en funksjon til å motta feilkode hvis heisen er fysisk forhindret
}
