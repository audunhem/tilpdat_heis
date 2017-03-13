
void timer_start(double door_open_duration);

void timer_stop();

bool timer_door_timeout();
//returns true when door should closed, constantly polled by main()
