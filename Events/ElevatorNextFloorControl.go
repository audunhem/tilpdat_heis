package Events

import (
  ."./../elevatorController"
  ."./../driver"
)


var thisElevator = Elevators[0]
//trenger noe ala thisElevator

/*func RecieveNewState(a,b,c,d chan int){
  select{
  case c <- x:
    //sett noe til noe
  case c2 <- y:
    //sett noe annet til noe annet
  }
}*/

func CheckIfShouldStop(elevatorData ElevatorData) bool {
  switch{
  case elevatorData.direction = DirnUp:
    if elevatorData.Orders[elevatorData.Floor][ButtonCallUp] == true || elevatorData.Orders[elevatorData.Floor][ButtonInternal] == true {
      //elevatorData.Orders[elevatorData.Floor][ButtonCallUp] = false
      //elevatorData.Orders[elevatorData.Floor][ButtonInternal] = false
      //mulig dette kan føre til at ordre forsvinner, og kanskje bedre med en egen funksjon for funksjonaliteten
      return true
    } else {
      if elevatorData.Floor == N_FLOORS - 1 {
        return true
      }
    } else if {
      for int i := elevatorData.Floor + 1 ; i < N_FLOORS ; i++ {
        if elevatorData.Orders[i][ButtonCallUp] == false || elevatorData.Orders[i][ButtonCallDown] == false || elevatorData.Orders[i][ButtonInternal] == false{
          return true
        }
      }
    }
    return false
  case elevatorData.direction = DirnDown:
    if elevatorData.Orders[elevatorData.Floor][ButtonCallDown] == true || elevatorData.Orders[elevatorData.Floor][ButtonInternal] == true {
      //elevatorData.Orders[elevatorData.Floor][ButtonCallUp] = false
      //elevatorData.Orders[elevatorData.Floor][ButtonInternal] = false
      //mulig dette kan føre til at ordre forsvinner, og kanskje bedre med en egen funksjon for funksjonaliteten
      return true
    } else if {
      if elevatorData.Floor == 0 {
        return true
      }
    } else if {
      for int i := 0 ; i < ElevatorData.Floor ; i++ {
        if elevatorData.Orders[i][ButtonCallUp] == false || elevatorData.Orders[i][ButtonCallDown] == false || elevatorData.Orders[i][ButtonInternal] == false{
          return true
        }
      }
    }
    return false
  }
}

/*func SendElevatorToNextFloor(newInternalOrder button){ //må vente til det kommer en intern ordre før man sender neste etasje
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
}*/





func OrderSetNextDirection(elevatorStruct ElevatorData) ElevatorData {
  elevatorData := elevatorStruct
  check := 0

  if elevatorData.Status == StatusIdle {
    for i := 0; i < N_FLOORS; i++ {
      for j := 0; j < N_BUTTONS; j++ {
        if elevatorData.Orders[i][j] == 1 {
          if elevatorData.Floor < i {
            elevatorData.Direction = DirnUp
            SetMotorDirection(DirnUp)
            elevatorData.Status = StatusMoving
          } else if elevatorData.Floor > i {
            elevatorData.Direction = DirnDown
            SetMotorDirection(DirnDown)
            elevatorData.Status = StatusMoving
          } else if elevatorData.Floor == i {
            elevatorData = fsmArriveAtFloor(i, elevatorData)
          }
        }

      }
    }

  } else if elevatorData.Direction == DirnUp {
    for i := elevatorData.Floor; i < N_FLOORS; i++ {
      for j := 0; j < N_BUTTONS; j++ {
        if elevatorData.Orders[i][j] == 1 {
          SetMotorDirection(DirnUp)
          check = 1
        }
      }
    }

    if check == 0 {
      for i := 0; i < elevatorData.Floor; i++ {
        for j := 0; j < N_BUTTONS; j++ {
          if elevatorData.Orders[i][j] == 1 {
            SetMotorDirection(DirnDown)
            elevatorData.Direction = DirnDown
          }
        }
      }
    } else if elevatorData.Direction == DirnDown {

      for i := 0; i < elevatorData.Floor; i++ {
        for j := 0; j < N_BUTTONS; j++ {
          if elevatorData.Orders[i][j] == 1 {
            SetMotorDirection(DirnDown)
            check = 1
          }
        }
      }

      if check == 0 {
        for i := elevatorData.Floor; i < N_FLOORS; i++ {
          for j := 0; j < N_BUTTONS; j++ {
            if elevatorData.Orders[i][j] == 1 {
              SetMotorDirection(DirnUp)
              elevatorData.Direction = DirnUp
            }
          }
        }
      }
    }
  }

  return elevatorData
}