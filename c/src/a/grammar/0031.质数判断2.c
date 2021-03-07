#include <stdbool.h>
#include <stdio.h>

bool isPrime(int n) {
  if (n < 0) { return false; }
  for (int i = 2; i < n; i++) {
    if (n % i == 0) { return false; }
  }
  return true;
}

int main(int argc, char *argv[]) {
  int num;
  printf("Please enter a number: ");
  scanf("%d", &num);
  bool res = isPrime(num);
  if (res) {
    printf("You input's %d is prime num\n", num);
  } else {
    printf("You input's %d is not a prime num\n", num);
  }
  return 0;
}
