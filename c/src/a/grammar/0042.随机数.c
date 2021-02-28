#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main(int argc, char *argv[]) {
  int a = rand();
  printf("固定的随机数：a = %d\n", a);

  int b;
  srand((unsigned) time(NULL));
  b = rand() % 10;
  printf("\n产生 0~9 之间的随机数：b = %d\n", b);

  int c;
  srand((unsigned) time(NULL));
  c = rand() % 51;
  printf("\n产生 0~50 之间的随机数：c = %d\n", c);

  int d;
  srand((unsigned) time(NULL));
  d = rand() % 101 + 14;
  printf("\n产生 14~114 之间的随机数：d = %d\n", d);
  return 0;
}
