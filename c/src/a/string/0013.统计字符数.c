#include <stdio.h>

int counter(const char *str);

int main(int argc, char *argv[]) {
  printf("counter = %d\n", counter("Hello World"));
  return 0;
}

int counter(const char *str) {
  int count = 0;
  while (*str) {
    putchar(*str++);
    count++;
  }
  putchar('\n');
  return count;
}
