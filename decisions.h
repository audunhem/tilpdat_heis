#include <stdbool.h>

void decision_reset_all_orders(struct ElevatorData* elevator);
//sets all orders to null

elev_motor_direction_t decision_next_motor_direction(struct ElevatorData* elevator);
//decides the direction the elevator is working in

void decision_remove_completed_orders(struct ElevatorData* elevator);

bool decision_check_if_should_stop(struct ElevatorData* elevator);
