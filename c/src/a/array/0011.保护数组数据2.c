#include <stdio.h>

int main(int argc, char *argv[]) {
  double rates[5] = {88.99, 77.12, 59.38, 11.29, 35.90};
  const double *ptr = rates;

  // error: read-only variable is not assignable
  // *ptr = 30.00;

  // error: read-only variable is not assignable
  // ptr[2] = 30.00;

  rates[0] = 99.99;
  printf("rate[0] = %.2f, *ptr = %.2f\n", rates[0], *ptr);

  ptr++;
  printf("rage[1] = %.2f, *ptr = %.2f\n", rates[1], *ptr);
  return 0;
}
