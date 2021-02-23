#include <stdarg.h>
#include <stdio.h>

double sum(int, ...);

int main(int argc, char *argv[]) {
  double s, t;
  s = sum(3, 1.1, 2.5, 13.3);
  t = sum(6, 1.1, 2.1, 13.1, 4.1, 5.1, 6.1);
  printf("Return value for sum(3, 1.1, 2.5, 13.3): %g\n", s);
  printf("Return value for sum(6, 1.1, 2.1, 13.1, 4.1, 5.1, 6.1): %g\n", t);
  return 0;
}

double sum(int lim, ...) {
  va_list ap; // 声明一个对象存储参数
  double total = 0;
  int i;
  va_start(ap, lim); // 把 ap 初始化为参数列表
  for (i = 0; i < lim; i++) {
    total += va_arg(ap, double); // 访问参数列表中的每一项
  }
  va_end(ap);
  return total;
}
