#include <stdio.h>

int getNumber() {
  return 999;
}

int main(int argc, char *argv[]) {
  int n = 99;

  const int a = getNumber();
  const int b = n;
  const float c = .1;

  printf("a = %d, b = %d, c = %f\n", a, b, c);
  return 0;
}
