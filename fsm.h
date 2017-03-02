#include "driver/elev.h"

void arrive_at_floor();
void external_button_pressed(int floor, int button);
void stop_at_floor();
void internal_button_pressed(int floor);
void set_lights(struct Elevator_data* elevator);
void initialize_elevator(struct Elevator_data* elevator);
