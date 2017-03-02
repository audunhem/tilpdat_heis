
#include "driver/elev.h"
#include "fsm.h"
#include "orders.h"
#include "timer.h"


int main(){
    struct Elevator_data elevator;
    elevator.current_floor = elev_get_floor_sensor_signal();
    elevator.direction = DIRN_STOP;
    initialize_elevator(&elevator);
  	//Lager variabel for å unngå å oppfatte tastetrykk flere ganger
  	int lastButtonPressed = -1;



  	while (1) {
  		//Vi ønsker kun beskjed hvis vi når en NY etasje! SKRIV DENNE PÅ EN BEDRE MÅTE, VI GJØR TRE KALL TIL GETFLOORSENSORSIGNAL
  		if (elev_get_floor_sensor_signal() != elevator.current_floor && elev_get_floor_sensor_signal() >= 0) {
  			elevator.current_floor = elev_get_floor_sensor_signal();
        arrive_at_floor(&elevator);
        start_timer(3.0);
        //trenger noe timer greie her før vi starter
        //leave_floor(&elevator);
  		}

  		//Looper gjennom alle EKSTERNE knapper
  		for (int i = 0; i < N_FLOORS; i++) {
  			for (int j = 0; j < 3; j++) {
  				if (elev_get_button_signal(j, i) == 1) {
  					if (lastButtonPressed != (N_FLOORS*i+j)) {
  						lastButtonPressed = N_FLOORS*i + j;
              struct Button_press order;
              order.floor = i;
              order.button = j;
              button_pressed(order, &elevator);
  					}

  				}
  			}
  		}

      if (door_timeout()){
        stop_timer();
        leave_floor(&elevator);
      }

      //trenger en form for timer

  		/*for (int i = 0; i < N_FLOORS; i++) {
  			if (elev_get_button_signal(2, i) == 1) {
  				if (lastButtonPressed != N_FLOORS*2+i) {
  					lastButtonPressed = N_FLOORS*2 + i;
            Button_press order;
            order.floor = i;
            order.button = j;
  					internal_button_pressed(order);
  				}
  			}
  		}*/
      set_lights(elevator.orders);
  	}
  return 0;
}
