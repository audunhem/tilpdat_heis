
void fsm_arrive_at_floor();
//is called every time the elevator enters a new floor

void fsm_order_button_pressed(struct ButtonPress order, struct ElevatorData* elevator);

void fsm_leave_floor(struct ElevatorData* elevator);
//is called if the elevator has stopped at a floor, and the door has closed

void fsm_stop_button_pressed(struct ElevatorData* elevator);

void fsm_initialize_elevator(struct ElevatorData* elevator);
//sets the elevator in a defined state
