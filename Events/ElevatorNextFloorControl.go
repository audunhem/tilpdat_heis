package main

import (
  "fmt"
  "sort"
)

const THIS_ELEVATOR = 1
const N_ELEVATORS = 3
const N_FLOORS = 4
const (
	DirnDown = -1 + iota
	DirnStop
	DirnUp
)

type Elevator struct{
  InternalOrders []int
  ExternalOrders []int //trenger noe for å vise retningen til ordren
  Direction int
  CurrentFloor int
  ID int
}

func RecieveNewState(a,b,c,d chan int){
  select{
  case c <- x:
    //sett noe til noe
  case c2 <- y:
    //sett noe annet til noe annet
  }
}

func SendElevatorToNextFloor(){
  //sender heisen til neste etasje i listen
}

func MotorOutOfOrder(){
  //trenger en funksjon til å motta feilkode hvis heisen er fysisk forhindret
}
