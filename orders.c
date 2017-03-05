#include "driver/elev.h"
#include <stdbool.h>
#include <stdio.h>

bool no_orders_above_floor(struct Elevator_data* elevator){
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

bool no_orders_below_floor(struct Elevator_data* elevator){
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

bool no_orders_at_current_floor(struct Elevator_data* elevator){
  if (elevator->orders[elevator->current_floor][BUTTON_CALL_UP] != 0 || elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] != 0 || elevator->orders[elevator->current_floor][BUTTON_COMMAND] != 0) {
    return false;
  }
  return true;
}

elev_motor_direction_t next_motor_direction(struct Elevator_data* elevator){
  switch (elevator->direction) {

  case (DIRN_UP):
    if (no_orders_above_floor(elevator)){
      return DIRN_DOWN;
    }
    elevator->direction = DIRN_UP;
    return DIRN_UP;

  case (DIRN_DOWN):
    if (no_orders_below_floor(elevator)){
      return DIRN_UP;
    }
    elevator->direction = DIRN_UP;
    return DIRN_DOWN;

  case (DIRN_STOP):
    for (int i = 0; i < N_BUTTONS; i++){
      if (elevator->orders[elevator->current_floor][i] == 1){
        return DIRN_STOP;
      }
    }
    if (no_orders_below_floor(elevator)){
      elevator->direction = DIRN_UP;
      return DIRN_UP;
    } else {
      elevator->direction = DIRN_DOWN;
      return DIRN_DOWN;
    }
  }

  if (no_orders_below_floor(elevator) && no_orders_above_floor(elevator) && no_orders_at_current_floor(elevator)){
    elevator->direction = DIRN_STOP;
    return DIRN_STOP;
  }
}

void remove_completed_orders(struct Elevator_data* elevator){
  switch (elevator->direction) {
  case (DIRN_UP):
    elevator->orders[elevator->current_floor][BUTTON_CALL_UP] = 0;
    elevator->orders[elevator->current_floor][BUTTON_COMMAND] = 0;

    if (no_orders_above_floor(elevator)) {
      elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] = 0;
    }

  case (DIRN_DOWN):
    elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] = 0;
    elevator->orders[elevator->current_floor][BUTTON_COMMAND] = 0;

    if (no_orders_below_floor(elevator)) {
      elevator->orders[elevator->current_floor][BUTTON_CALL_UP] = 0;
    }

  case (DIRN_STOP):
    printf("Error in remove_completed_orders");
  }
}

bool check_if_should_stop(struct Elevator_data* elevator) {
  switch (elevator->direction) {

  case (DIRN_UP):
    if (elevator->orders[elevator->current_floor][BUTTON_CALL_UP] == 1 || elevator->orders[elevator->current_floor][BUTTON_COMMAND] == 1) {
      return true;
    } else if (no_orders_above_floor(elevator)){
      return true;
    }

  case (DIRN_DOWN):
    if (elevator->orders[elevator->current_floor][BUTTON_CALL_DOWN] == 1 || elevator->orders[elevator->current_floor][BUTTON_COMMAND] == 1) {
      return true;
    } else if (no_orders_below_floor(elevator)){
      return true;
    }
  case (DIRN_STOP):
    printf("Error in check_if_should_stop");
  }
  return false;
}
