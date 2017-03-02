#include <time.h>
#include <stdbool.h>



static time_t start_time;
static double duration;
static bool timer_active;

void start_timer(double door_open_duration){
  duration = door_open_duration;
  start_time = time(NULL);
  timer_active = true;
}

void stop_timer(){
  timer_active = false;
}

bool door_timeout(){
  if (timer_active && difftime(time(NULL), start_time) > duration){
    return true;
  }
  return false;
}
