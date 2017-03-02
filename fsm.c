#include "driver/elev.h"
#include "orders.h"
#include "fsm.h"
#include <stdio.h>

void arrive_at_floor(struct Elevator_data* elevator){
  elev_set_floor_indicator(elevator.current_floor);
	if (check_if_should_stop(&elevator) == true) {
    elev_set_motor_direction(DIRN_STOP);
  	elev_set_door_open_lamp(1);
    //start en timer
	}
}

void button_pressed(struct Button_press order, struct Elevator_data* elevator){
	elevator.order[order.floor][order.button] = 1;
	if elevator.direction == DIRN_STOP {
		printf("IDLE");
		set_next_direction(&elevator);
	}
  set_lights(&elevator);
}

void leave_floor(struct Elevator_data* elevator){
  remove_completed_orders(&elevator);
  set_lights(&elevator.orders);
  elev_set_motor_direction(next_motor_direction(&elevator));
}

void set_lights(struct Elevator_data* elevator){
  for (int i = 0; i < N_FLOORS; i++){
    for (int j = 0; j < N_BUTTONS; j++){
      elev_set_button_lamp(j, i, elevator.orders[i][j]);
    }
  }
}
void initialize_elevator(struct elevator){
  elev_init();

	if (elev_get_floor_sensor_signal() == -1) {
		elev_set_motor_direction(DIRN_UP);
		while (elev_get_floor_sensor_signal() == -1){
      elev_set_motor_direction(DIRN_STOP);
    }
	}
}
