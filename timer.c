#include <time.h>
#include <stdbool.h>



static time_t start_time;
static double duration;
static bool timer_active;

void timer_start(double DOOR_OPEN_DURATION){
  duration = DOOR_OPEN_DURATION;
  start_time = time(NULL);
  timer_active = true;
}

void timer_stop(){
  timer_active = false;
}

bool timer_door_timeout(){

  if (timer_active && difftime(time(NULL), start_time) > duration){
    return true;
  }

  return false;
}
