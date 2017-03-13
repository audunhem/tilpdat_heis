
#include "driver/elev.h"
#include "fsm.h"
#include "orders.h"
#include "timer.h"
#include <stdlib.h>
#include <stdio.h>
void print_orders(struct Elevator_data elevator){
  for (int i = 0; i < N_FLOORS; i++){
    for (int j = 0; j < N_BUTTONS; j++){
      printf("%i",elevator.orders[i][j]);
    }
    printf("\n");
  }
  printf("Elevator direction: ");
  printf("%i\n",elevator.direction);
  printf("Elevator floor: ");
  printf("%i\n",elevator.current_floor);
  printf("--------------------------\n");

}

int main(){
    struct Elevator_data elevator;
    //all data about the elevator, it's being passed around to other functions
    fsm_initialize_elevator(&elevator);
    //to stop one button pressed from being added several times
  	int lastButtonPressed = -1;

  	while (true) {

      //checks if elevator has arrived at new floor
  		if (elev_get_floor_sensor_signal() != elevator.current_floor && elev_get_floor_sensor_signal() >= 0) {
  			elevator.current_floor = elev_get_floor_sensor_signal();
        fsm_arrive_at_floor(&elevator);
        print_orders(elevator);
  		}

      //checks if elevator should leave current floor
      if (timer_door_timeout()){
        timer_stop();
        if (elev_get_floor_sensor_signal() >= 0) {
          lastButtonPressed = -1;
          fsm_leave_floor(&elevator);
        }
      }

      //looping through all order buttons
  		for (int floor = 0; floor < N_FLOORS; floor++) {
  			for (int button = BUTTON_CALL_UP; button < N_BUTTONS; button++) {
          if (elev_get_button_signal(button, floor) == 1 && lastButtonPressed != (N_FLOORS*floor + button) && elevator.orders[floor][button] == 0) {
  					lastButtonPressed = N_FLOORS*floor + button;
            struct Button_press order;
            order.floor = floor;
            order.button = button;
            fsm_order_button_pressed(order, &elevator);
  				  print_orders(elevator);
          }
  			}
  		}

      //checks if stop button is pressed
      if (elev_get_stop_signal()){
        fsm_stop_button_pressed(&elevator);
      }
  	}
  return 0;
}
