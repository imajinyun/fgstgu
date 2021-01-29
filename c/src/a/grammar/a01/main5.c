#include <stdio.h>
#include <math.h>

int main(int argc, char *argv[]) {
  char char1 = 127;
  short short1 = pow(2, 8);
  int int1 = pow(2, 16);
  long long1 = pow(2, 32);
  float float1 = 9.99e10;
  double double1 = 9.99e20;

  printf("char takes %lu bytes, char1 = %c\n", sizeof(char1), char1);
  printf("short takes %lu bytes, short1 = %hd\n", sizeof(short1), short1);
  printf("int takes %lu bytes, int1 = %d\n", sizeof(int1), int1);
  printf("long takes %lu bytes, long1 = %ld\n", sizeof(long1), long1);
  printf("float takes %lu bytes, float1 = %f\n", sizeof(float1), float1);
  printf("double takes %lu bytes, double1 = %f\n", sizeof(double1), double1);
  return 0;
}
