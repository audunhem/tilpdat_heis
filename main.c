
#include "driver/elev.h"
#include "fsm.h"


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
  		}

  		//Looper gjennom alle EKSTERNE knapper
  		for (int i = 0; i < N_FLOORS; i++) {
  			for (int j = 0; j < 2; j++) {
  				if (elev_get_button_signal(j, i) == 1) {
  					if (lastButtonPressed != (2*i+j)) {
  						lastButtonPressed = 2*i + j;
              external_button_pressed(i,j);
  					}

  				}
  			}
  		}


  		for (int i = 0; i < N_FLOORS; i++) {
  			if (elev_get_button_signal(2, i) == 1) {
  				if (lastButtonPressed != N_FLOORS*2+i) {
  					lastButtonPressed = N_FLOORS*2 + i;
  					internal_button_pressed(i);
  				}
  			}
  		}
      set_lights(&elevator);
  	}
  return 0;
}
