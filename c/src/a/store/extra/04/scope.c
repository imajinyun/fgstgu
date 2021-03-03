#include <stdio.h>
#include "scope.h"
#include "../../../hello/debug/dbg.h"

int THE_SIZE = 1000;

static int THE_AGE = 49;

int getAge() {
  return THE_AGE;
}

void setAge(int age) {
  THE_AGE = age;
}

double updateRatio(double new_ratio) {
  static double ratio = 1.0;
  double old_ratio = ratio;
  ratio = new_ratio;
  return old_ratio;
}

void printSize() {
  info("I think size is %d", THE_SIZE);
}
