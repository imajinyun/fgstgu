#include <stdio.h>

long factorial(int n) {
  int i;
  long res = 1;
  for (i = 1; i <= n; i++) { res *= i; }
  return res;
}

long sum(int n) {
  int i;
  long res = 0;
  for (i = 0; i <= n; i++) { res += factorial(i); }
  return res;
}

int main(int argc, char *argv[]) {
  printf("阶乘求和：1! + 2! + ... + 9! + 10! = %ld\n", sum(10));
  return 0;
}
