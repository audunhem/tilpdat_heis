
void arrive_at_floor();

void button_pressed(struct Button_press order, struct Elevator_data* elevator);

void leave_floor(struct Elevator_data* elevator);

void set_lights(int orders[N_FLOORS][N_BUTTONS]);

void initialize_elevator(struct Elevator_data* elevator);
