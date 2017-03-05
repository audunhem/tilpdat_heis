
void arrive_at_floor();
//is called every time the elevator enters a new floor

void order_button_pressed(struct Button_press order, struct Elevator_data* elevator);

void leave_floor(struct Elevator_data* elevator);
//is called if the elevator has stopped at a floor, and the door has closed

void set_lights(int orders[N_FLOORS][N_BUTTONS]);
//sets all lights in the system

void initialize_elevator(struct Elevator_data* elevator);
//sets the elevator in a defined state

void stop_button_pressed(struct Elevator_data* elevator);
