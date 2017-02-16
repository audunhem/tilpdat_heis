package Events

import (

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
    case newInternalOrder.floor > thisElevator.currentFloor:
      
  }
}

func MotorOutOfOrder(){
  //trenger en funksjon til å motta feilkode hvis heisen er fysisk forhindret
}
