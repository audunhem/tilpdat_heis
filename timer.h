
void start_timer(double door_open_duration);

void stop_timer();

bool door_timeout();
//returns true when door should closed, constantly polled by main()
