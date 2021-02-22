#include <stdio.h>

int main(int argc, char *argv[]) {
  double d1;
  char c1;
  char c2;
  double d2;
  char c3;
  char _Alignas(double) c4;

  // double 的对齐值是 8，这意味着地址的类型对齐可以被 8 整除。
  // 以 0 或 8 结尾的十六进制地址可被 8 整除。
  // 这就是地址用两个 double 类型的变量（d1 和 d2）和 char 类型的变量 c4（该变量是 double 对齐值）。
  printf("char alignment: %zd\n", _Alignof(char));
  printf("double alignment: %zd\n", _Alignof(double));
  printf("&d1: %p\n", &d1);
  printf("&c1: %p\n", &c1);
  printf("&c2: %p\n", &c2);
  printf("&d2: %p\n", &d2);
  printf("&c3: %p\n", &c3);
  printf("&c4: %p\n", &c4);
  return 0;
}
