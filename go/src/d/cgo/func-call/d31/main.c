#include <stdio.h>

void goPrintln(char *);
int average(int a, int b);

// -> gcc -o main.o main.c main.a
// -> ./main.o
int main(int argc, char *argv[]) {
  int a, b, x;
  a = 10;
  b = 30;
  x = average(a, b);
  printf("(%d+%d)=%d\n", a, b, x);
  goPrintln("Done...");
  return 0;
}
