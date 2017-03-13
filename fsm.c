#include "driver/elev.h"
#include "orders.h"
#include "fsm.h"
#include "timer.h"
#include <stdio.h>

static void fsm_set_lights(int orders[N_FLOORS][N_BUTTONS]){

  for (int i = 0; i < N_FLOORS; i++){
    for (int j = 0; j < N_BUTTONS; j++){
      elev_set_button_lamp(j, i, orders[i][j]);
    }
  }
}

void fsm_arrive_at_floor(struct ElevatorData* elevator){

  elev_set_floor_indicator(elevator->current_floor);

	if (logic_check_if_should_stop(elevator)) {
    elev_set_motor_direction(DIRN_STOP);
  	elev_set_door_open_lamp(1);
    timer_start(DOOR_OPEN_DURATION);
	}
}

void fsm_order_button_pressed(struct Button_press order, struct ElevatorData* elevator){

	elevator->orders[order.floor][order.button] = 1;
  	fsm_set_lights(elevator->orders);

	if ((elevator->direction == DIRN_STOP && !elev_get_door_open_lamp()) || elevator->stopped_between_floors) {
    elev_set_motor_direction(logic_next_motor_direction(elevator));
    elevator->stopped_between_floors = false;

    if (logic_next_motor_direction(elevator) == DIRN_STOP) {
      //order is called at current floor
      fsm_arrive_at_floor(elevator);
    }
	}
}

void fsm_stop_button_pressed(struct ElevatorData* elevator){

  elev_set_motor_direction(DIRN_STOP);
  timer_stop();
  logic_reset_all_orders(elevator);
  fsm_set_lights(elevator->orders);

  if (elev_get_floor_sensor_signal() >= 0){
    elev_set_door_open_lamp(1);
  } else {
    elevator->stopped_between_floors = true;
  }

  while (elev_get_stop_signal()){
    elev_set_stop_lamp(1);
  }

  if (elev_get_floor_sensor_signal() >= 0){
    timer_start(DOOR_OPEN_DURATION);
  }

  elev_set_stop_lamp(0);
}

void fsm_leave_floor(struct ElevatorData* elevator){
	//logic_next_motor_direction(elevator);
  elev_set_door_open_lamp(0);
  logic_remove_completed_orders(elevator);
  fsm_set_lights(elevator->orders);
  elev_set_motor_direction(logic_next_motor_direction(elevator));
}


void fsm_initialize_elevator(struct ElevatorData* elevator){

  bool initialized = false;
  bool stopped_between_floors = false;
  initialized = elev_init();

  if (!initialized){
    printf("Initializing failed!");
  }

  logic_reset_all_orders(elevator);

	if (elev_get_floor_sensor_signal() == -1) {
		elev_set_motor_direction(DIRN_UP);
    elevator->direction = DIRN_UP;

		while (elev_get_floor_sensor_signal() == -1){
      //do nothing
    }
	}

  elev_set_motor_direction(DIRN_STOP);
  elevator->direction = DIRN_STOP;
  elevator->current_floor = elev_get_floor_sensor_signal();
}
