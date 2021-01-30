#include <stdio.h>

int main(int argc, char *argv[]) {
  puts("浮点数的长度是固定的，float 始终占用 4 个字节，double 始终占用 8 个字节。\n");

  float a = 0.1234;
  float b = 1234.1234;
  float c = 1.2345E3;
  float d = 1.123456789;
  double e = 1234;
  double f = 1.2345e-4;
  double g = 1.123456789;

  puts(
      "一般的输出方式：\n"
      "1. %f 和 %lf 默认保留六位小数，不足六位以 0 补齐，超过六位按四舍五入截断；"
      "2. 将整数赋值给 float 变量时会变成小数；"
      "3. 以指数形式输出小数时，输出结果为科学计数法；也就是说，尾数部分的取值为：0 ≤ 尾数 < 10；"
  );
  printf("a = %e\n", a); // 1.234000e-01
  printf("b = %f\n", b); // 1234.123413
  printf("c = %f\n", c); // 1234.500000
  printf("d = %lE\n", d); // 1.123457E+00
  printf("e = %lf\n", e); // 1234.000000
  printf("f = %f\n", f); // 0.000123
  printf("g = %lf\n", g); // 1.123457

  float f1 = 0.00001;
  float f2 = 40000000;
  float f3 = 13.33;
  float f4 = 1.123456789;

  puts("\n智能的输出方式，使用 %g：");
  printf("f1 = %g\n", f1);
  printf("f2 = %g\n", f2);
  printf("f3 = %g\n", f3);
  printf("f4 = %g\n", f4);

  short x = 32L;
  int y = 300;
  long z = 100l;

  float l = 52.55f;
  float m = 19.6F;
  float n = 0.123;

  short o = 56.5678;
  int p = 100l;
  long q = 1.5689E2;

  puts("\n浮点数的其它表示方法（以下输出中有丢失精度的错误）：");
  printf("x = %f x = %d\n", x, x); // x = 0.000000 x = 32
  printf("y = %f y = %d\n", y, y); // y = 0.000000 y = 300
  printf("z = %f z = %d\n", z, z); // z = 0.000000 z = 100
  printf("l = %f l = %d\n", l, l); // l = 52.549999 l = 73896
  printf("m = %f m = %d\n", m, m); // m = 19.600000 m = 73896
  printf("n = %f n = %d\n", n, n); // n = 0.123000 n = 73896
  printf("o = %f o = %d\n", o, o); // o = 0.123000 o = 56
  printf("p = %f p = %d\n", p, p); // p = 0.123000 p = 100
  printf("q = %f q = %d\n", q, q); // q = 0.123000 q = 156

  return 0;
}
