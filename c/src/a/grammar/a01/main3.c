#include <stdio.h>

// 基本运算。
int main(int argc, char *argv[]) {
  int a = 100, b = 11;

  puts("基本运算：\n");
  puts("一旦除数和被除数中有一个是小数，那么运算结果也是小数，并且是 double 类型的小数。\n");

  printf("a + b = %d\n", a + b);
  printf("a - b = %d\n", a - b);
  printf("a * b = %d\n", a * b);
  printf("a / b = %d\n", a / b);
  printf("a %% b = %d\n", a % b);
  printf("a / c = %f\n", a / 10.0); // 10.000000

  puts("\n自增自减：\n");

  int i, j, m, n;
  i = 1, 2, 3, 4, 5; // i = 1
  j = (1, 2, 3, 4, 5); // j = 5
  m = ++i, ++i, i++, i++;
  n = (++j, ++j, ++j, ++j);
  printf("i = %d\n", i); // 5
  printf("j = %d\n", j); // 9
  printf("m = %d\n", m); // 2
  printf("m = %d\n", n); // 9
  return 0;
}
