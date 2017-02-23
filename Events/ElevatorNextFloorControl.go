package Events

import (
  . "./../driver"
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
  switch {
  case elevatorData.Direction == DirnUp:
    if elevatorData.Orders[elevatorData.Floor][ButtonCallUp] == 1 || elevatorData.Orders[elevatorData.Floor][ButtonInternal] == 1 {
      return true
    } else if elevatorData.Floor == N_FLOORS-1 {
      return true

    } else {
      for i := elevatorData.Floor + 1; i < N_FLOORS; i++ {
        if elevatorData.Orders[i][ButtonCallUp] != 0 || elevatorData.Orders[i][ButtonCallDown] != 0 || elevatorData.Orders[i][ButtonInternal] != 0 {
          return false
        }
      }
      return true
    }
    return false
  case elevatorData.Direction == DirnDown:
    if elevatorData.Orders[elevatorData.Floor][ButtonCallDown] == 1 || elevatorData.Orders[elevatorData.Floor][ButtonInternal] == 1 {
      //elevatorData.Orders[elevatorData.Floor][ButtonCallUp] = false
      //elevatorData.Orders[elevatorData.Floor][ButtonInternal] = false
      //mulig dette kan føre til at ordre forsvinner, og kanskje bedre med en egen funksjon for funksjonaliteten
      return true
    } else if elevatorData.Floor == 0 {
      return true
    } else {
      for i := 0; i < elevatorData.Floor; i++ {
        if elevatorData.Orders[i][ButtonCallUp] != 0 || elevatorData.Orders[i][ButtonCallDown] != 0 || elevatorData.Orders[i][ButtonInternal] != 0 {
          return false
        }
      }
      return true
    }
    return false
  }
  return false
}

func OrderCompleted(elevatorData ElevatorData) ElevatorData {
  switch elevatorData.Direction {
  case DirnUp:
    if elevatorData.Orders[elevatorData.Floor][ButtonCallUp] == 1 || elevatorData.Orders[elevatorData.Floor][ButtonInternal] == 1 {
      elevatorData.Orders[elevatorData.Floor][ButtonCallUp] = 0
      elevatorData.Orders[elevatorData.Floor][ButtonInternal] = 0
    } else if elevatorData.Floor == N_FLOORS-1 {
      elevatorData.Orders[elevatorData.Floor][ButtonCallDown] = 0
      elevatorData.Orders[elevatorData.Floor][ButtonInternal] = 0
    } else {
      for i := elevatorData.Floor + 1; i < N_FLOORS; i++ {
        if elevatorData.Orders[i][ButtonCallUp] == 0 || elevatorData.Orders[i][ButtonCallDown] == 0 || elevatorData.Orders[i][ButtonInternal] == 0 {
          elevatorData.Orders[elevatorData.Floor][ButtonCallDown] = 0
        }
      }
    }
  case DirnDown:
    if elevatorData.Orders[elevatorData.Floor][ButtonCallDown] == 1 || elevatorData.Orders[elevatorData.Floor][ButtonInternal] == 1 {
      elevatorData.Orders[elevatorData.Floor][ButtonCallDown] = 0
      elevatorData.Orders[elevatorData.Floor][ButtonInternal] = 0
    } else if elevatorData.Floor == 0 {
      elevatorData.Orders[elevatorData.Floor][ButtonCallUp] = 0
      elevatorData.Orders[elevatorData.Floor][ButtonInternal] = 0
    } else {
      for i := 0; i < elevatorData.Floor; i++ {
        if elevatorData.Orders[i][ButtonCallUp] == 0 || elevatorData.Orders[i][ButtonCallDown] == 0 || elevatorData.Orders[i][ButtonInternal] == 0 {
          elevatorData.Orders[elevatorData.Floor][ButtonCallUp] = 0
        }
      }
    }
  }
  return elevatorData
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
            //FIKS DETTE
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
            check = 1
          }
        }
      }

      if check == 0 {
        elevatorData.Status = StatusIdle
        elevatorData.Direction = DirnStop
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
              check = 1
            }
          }
        }
      }

      if check == 0 {
        check = 1
      }
    }

  } else {
    elevatorData.Status = StatusIdle
    elevatorData.Direction = DirnStop
  }

  return elevatorData
}
