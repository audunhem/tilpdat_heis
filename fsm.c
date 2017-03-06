#include "driver/elev.h"
#include "orders.h"
#include "fsm.h"
#include "timer.h"
#include <stdio.h>

static void set_lights(int orders[N_FLOORS][N_BUTTONS]){

  for (int i = 0; i < N_FLOORS; i++){
    for (int j = 0; j < N_BUTTONS; j++){
      elev_set_button_lamp(j, i, orders[i][j]);
    }
  }
}

void arrive_at_floor(struct Elevator_data* elevator){

  elev_set_floor_indicator(elevator->current_floor);

	if (check_if_should_stop(elevator)) {
    elev_set_motor_direction(DIRN_STOP);
  	elev_set_door_open_lamp(1);
    start_timer(DOOR_OPEN_DURATION);
	}
}

void order_button_pressed(struct Button_press order, struct Elevator_data* elevator){

	elevator->orders[order.floor][order.button] = 1;
  set_lights(elevator->orders);

	if (elevator->direction == DIRN_STOP && !elev_get_door_open_lamp()) {
    elev_set_motor_direction(next_motor_direction(elevator));

    if (next_motor_direction(elevator) == DIRN_STOP) {
      arrive_at_floor(elevator);
    }
	}
}

void stop_button_pressed(struct Elevator_data* elevator){

  elev_set_motor_direction(DIRN_STOP);
  elevator->direction = DIRN_STOP;
  stop_timer();

  for (int i = 0; i < N_FLOORS; i++){
    for (int j = 0; j < N_BUTTONS; j++){
      elevator->orders[i][j] = 0;
    }
  }
  set_lights(elevator->orders);

  if (elev_get_floor_sensor_signal() >= 0){
    elev_set_door_open_lamp(1);
  }

  while (elev_get_stop_signal()){
    elev_set_stop_lamp(1);
  }

  if (elev_get_floor_sensor_signal() >= 0){
    start_timer(DOOR_OPEN_DURATION);
  }

  elev_set_stop_lamp(0);
}

void leave_floor(struct Elevator_data* elevator){

  remove_completed_orders(elevator);
  set_lights(elevator->orders);
  elev_set_door_open_lamp(0);
  elev_set_motor_direction(next_motor_direction(elevator));
}


void initialize_elevator(struct Elevator_data* elevator){

  bool initialized = false;
  initialized = elev_init();

  if (!initialized){
    printf("Initializing failed!");
  }

  for (int i = 0; i < N_FLOORS; i++){
    for (int j = 0; j < N_BUTTONS; j++){
      elevator->orders[i][j] = 0;
    }
  }

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
