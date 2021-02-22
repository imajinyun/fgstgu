#include <stdio.h>

#define SQUARE(x) x *x
#define SQUARE2(x) ((x) * (x))
#define PR(x) printf("The result is %d.\n", x)

int main(int argc, char *argv[]) {
  int x = 5, z;
  printf("x = %d\n", x);
  z = SQUARE(x);
  printf("Evaluating SQUARE(x): ");
  PR(z); // 25

  z = SQUARE(2);
  printf("Evaluating SQUARE(2): ");
  PR(z); // 4

  printf("\n");

  // x*x 变成了 x+2*x+2。如果 x 为 5，那么该表达式的值为：5+2*5+2 = 5+10+2 = 17
  printf("Evaluating SQUARE(x+2): ");
  PR(SQUARE(x + 2)); // 17
  printf("Evaluating SQUARE(x+2): ");
  PR(SQUARE2(x + 2)); // 49

  // x*x 变成了 2*2。如果 x 为 5，那么该表达式的值为：100/2*2 = 50*2 = 100
  printf("Evaluating 100/SQUARE(2): ");
  PR(100 / SQUARE(2)); // 100
  printf("Evaluating 100/SQUARE(2): ");
  PR(100 / SQUARE2(2)); // 25

  // x*x 变成了 ++x*++x。如果 x 为 5，那么该表达式的值为：6*7 = 42
  printf("x is %d.\n", x); // 5
  printf("Evaluating SQUARE(++x): ");
  PR(SQUARE(++x)); // 42

  // x*x 变成了 (++x)*(++x)。程序执行到这儿时，x 的值为 7，那么该表达式的值为：8*9 = 72
  printf("Evaluating SQUARE(++x): ");
  PR(SQUARE2(++x)); // 72

  printf("\n");

  printf("After incrementing, x is %x.\n", x); // 9
  return 0;
}
