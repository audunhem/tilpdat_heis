#include <stdbool.h>

elev_motor_direction_t next_motor_direction(struct Elevator_data* elevator);

void remove_completed_orders(struct Elevator_data* elevator);

bool check_if_should_stop(struct Elevator_data* elevator);
