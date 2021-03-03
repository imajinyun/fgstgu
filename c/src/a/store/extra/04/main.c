#include <stdio.h>
#include "scope.h"
#include "../../../hello/debug/dbg.h"

const char *MY_NAME = "John H Zend";

void testScope(int n) {
  info("n value: %d", n);
  if (n > 10) {
    int n = 200;
    info("n in this scope: %d", n);
  }
  info("n is at exit: %d", n);
  n = 3000;
  info("n after assign: %d", n);
}

int main(int argc, char *argv[]) {
  info("My name: %s, age: %d", MY_NAME, getAge());
  setAge(99);
  info("My age is now: %d", getAge());

  info("THE_SIZE value: %d", THE_SIZE);
  printSize();

  THE_SIZE = 999;
  printSize();

  info("Ratio at first: %f", updateRatio(2.0));
  info("Ratio again: %f", updateRatio(5.0));
  info("Ratio once more: %f", updateRatio(9.0));

  // test the scope demo
  int n = 5;
  testScope(n);
  testScope(n * 20);
  info("n after calling testScope(n * 20): %d", n);
  return 0;
}
