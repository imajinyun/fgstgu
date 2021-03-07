#include <stdio.h>
#include <stdlib.h>

long factorial(int num);

int main(int argc, char *argv[]) {
  int num;
  printf("Please enter a number: ");
  scanf("%d", &num);
  if (num < 0 || num > 15) {
    printf("Please enter between 0 and 15 number\n");
    exit(EXIT_FAILURE);
  }
  long res = factorial(num);
  printf("%d factorial result is %ld\n", num, res);
  return 0;
}

long factorial(int num) {
  long res;
  if (num == 0 || num == 1) {
    return 1;
  } else {
    res = num * factorial(num - 1);
  }
  return res;
}
