#include <stdio.h>

#define MAX(x, y) ((x) > (y) ? (x) : (y))
#define MIN(x, y) ((x) < (y) ? (x) : (y))
#define ABS(x) ((x) < 0 ? -(x) : (x))
#define ISSIGN(x) ((x) == '+' || (x) == '-' ? 1 : 0)

int main(int argc, char *argv[]) {
  int x = 3, y = 5, z = -10;
  printf("mas value is %d\n", MAX(x, y));
  printf("min value is %d\n", MIN(x, y));
  printf("abs value is %d\n", ABS(z));
  printf("issign value is %d\n", ISSIGN('-'));
  return 0;
}
