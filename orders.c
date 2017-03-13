#include "driver/elev.h"
#include <stdbool.h>
#include <stdio.h>

void logic_reset_all_orders(struct ElevatorData* elevator){
  for (int i = 0; i < N_FLOORS; i++){
    for (int j = 0; j < N_BUTTONS; j++){
      elevator->orders[i][j] = 0;
    }
  }
}

static bool no_orders_below_floor(struct ElevatorData* elevator){

  if (elevator->current_floor == 0) {
    return true;
  }

  for (int i = 0; i < elevator->current_floor; i++) {
    if (elevator->orders[i][BUTTON_CALL_UP] != 0 || elevator->orders[i][BUTTON_CALL_DOWN] != 0 || elevator->orders[i][BUTTON_COMMAND] != 0) {
      return false;
    }
  }

  return true;
}

static bool no_orders_above_floor(struct ElevatorData* elevator){

  if (elevator->current_floor == N_FLOORS-1) {
    return true;
  }

  for (int i = elevator->current_floor + 1; i < N_FLOORS; i++) {
    if (elevator->orders[i][BUTTON_CALL_UP] != 0 || elevator->orders[i][BUTTON_CALL_DOWN] != 0 || elevator->orders[i][BUTTON_COMMAND] != 0) {
      return false;
    }
  }

  return true;
}

static bool no_orders_at_current_floor(struct ElevatorData* elevator){

  if (elevator->orders[elevator->current_floor][BUTTON_CALL_UP] != 0 || elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] != 0 || elevator->orders[elevator->current_floor][BUTTON_COMMAND] != 0) {
    return false;
  }

  return true;
}

//problemer når heisen får en ordre i samme etasje når den har stått stille
//noen knapper reagerer ikke alltid

elev_motor_direction_t logic_next_motor_direction(struct ElevatorData* elevator){

  if (elev_get_floor_sensor_signal() != -1){
    //if the elevator is at a floor
  if (no_orders_below_floor(elevator) && no_orders_above_floor(elevator) && no_orders_at_current_floor(elevator)){
    elevator->direction = DIRN_STOP;
    return DIRN_STOP;
  }

  switch (elevator->direction) {

  case (DIRN_UP):

    if (no_orders_above_floor(elevator)){
      elevator->direction = DIRN_DOWN;
      return DIRN_DOWN;
    }

    elevator->direction = DIRN_UP;
    return DIRN_UP;

  case (DIRN_DOWN):

    if (no_orders_below_floor(elevator)){
      elevator->direction = DIRN_UP;
      return DIRN_UP;
    }

    elevator->direction = DIRN_DOWN;
    return DIRN_DOWN;

  case (DIRN_STOP):

    if(!no_orders_at_current_floor(elevator)){
      //to stop the elevator from picking up all orders at a floor
      if (!no_orders_below_floor(elevator)){
        elevator->direction = DIRN_DOWN;
        return DIRN_DOWN;
      } else if (!no_orders_above_floor(elevator)){
        elevator->direction = DIRN_UP;
        return DIRN_UP;
      }
      elevator->direction = DIRN_STOP;
      return DIRN_STOP;
    }

    if (no_orders_below_floor(elevator)){
      elevator->direction = DIRN_UP;
      return DIRN_UP;
    } else {
      elevator->direction = DIRN_DOWN;
      return DIRN_DOWN;
    }

    break;
  }

  }

  //when the elevator is standing still between floors
  switch (elevator->direction){
    case DIRN_UP:
      if (no_orders_above_floor(elevator)){
        //incrementing to give regular behaviour when switching direction
        elevator->current_floor++;
        elevator->direction = DIRN_DOWN;
        return DIRN_DOWN;
      } else {
        elevator->direction = DIRN_UP;
        return DIRN_UP;
      }
    case DIRN_DOWN:
      if (no_orders_below_floor(elevator)){
        //decrementing to give regular behaviour when switching direction
        elevator->current_floor--;
        elevator->direction = DIRN_UP;
        return DIRN_UP;
      } else {
        elevator->direction = DIRN_DOWN;
        return DIRN_DOWN;
    }
  }

  return DIRN_STOP;

}

void logic_remove_completed_orders(struct ElevatorData* elevator){

  switch (logic_next_motor_direction(elevator)) {

  case (DIRN_UP):

    elevator->orders[elevator->current_floor][BUTTON_CALL_UP] = 0;
    elevator->orders[elevator->current_floor][BUTTON_COMMAND] = 0;
    if (no_orders_above_floor(elevator)) {
      elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] = 0;
    }
    break;

  case (DIRN_DOWN):

    elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] = 0;
    elevator->orders[elevator->current_floor][BUTTON_COMMAND] = 0;

    if (no_orders_below_floor(elevator)) {
      elevator->orders[elevator->current_floor][BUTTON_CALL_UP] = 0;
    }
    break;

  case (DIRN_STOP):
    if (!no_orders_below_floor(elevator)) {
      elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] = 0;
    } else if (!no_orders_above_floor(elevator)) {
      elevator->orders[elevator->current_floor][BUTTON_CALL_UP] = 0;
    } else if (no_orders_above_floor(elevator) && no_orders_below_floor(elevator)){
      elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] = 0;
      elevator->orders[elevator->current_floor][BUTTON_CALL_UP] = 0;
    } else if (elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN]*elevator->orders[elevator->current_floor][BUTTON_CALL_UP] == 0){
      elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] = 0;
      elevator->orders[elevator->current_floor][BUTTON_CALL_UP] = 0;
    }
    elevator->orders[elevator->current_floor][BUTTON_COMMAND] = 0;

    break;
  }
}

bool logic_check_if_should_stop(struct ElevatorData* elevator) {

  switch (elevator->direction) {

  case (DIRN_UP):

    if (elevator->orders[elevator->current_floor][BUTTON_CALL_UP] == 1 || elevator->orders[elevator->current_floor][BUTTON_COMMAND] == 1) {
      return true;
    } else if (no_orders_above_floor(elevator)){
      return true;
    }
    break;

  case (DIRN_DOWN):

    if (elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] == 1 || elevator->orders[elevator->current_floor][BUTTON_COMMAND] == 1) {
      return true;
    } else if (no_orders_below_floor(elevator)){
      return true;
    }
    break;

  case (DIRN_STOP):
    return true;
  }
  return false;
}
