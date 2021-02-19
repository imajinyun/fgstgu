#include <stdio.h>

#define SIZE 10

int sum(int *start, int *end);

int main(int argc, char *argv[]) {
  int marbles[SIZE] = {20, 10, 5, 39, 4, 16, 19, 26, 31, 20};
  long res;
  res = sum(marbles, marbles + SIZE);
  printf("The total number of marbles is %ld\n", res);
  return 0;
}

int sum(int *start, int *end) {
  int total = 0;
  while (start < end) {
    total += *start;
    start++;
  }
  return total;
}
