#include <assert.h>
#include <math.h>
#include <stdio.h>

#define NDEBUG

int main(int argc, char *argv[]) {
  double x, y, z;
  puts("Enter a pair of numbers (0 0 to quit):");
  while (scanf("%lf%lf", &x, &y) == 2 && (x != 0 || y != 0)) {
    z = x * x - y * y;
    assert(z >= 0);
    printf("Answer is %f\n", sqrt(z));
    puts("Next pair of numbers:");
  }
  puts("Done!\n");
  return 0;
}
