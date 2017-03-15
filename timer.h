
void timer_start(double countdown_time);

void timer_stop();

bool timer_door_timeout();
//returns true when door should close, constantly polled by main()
