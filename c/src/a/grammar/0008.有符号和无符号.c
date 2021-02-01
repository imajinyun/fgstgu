#include <stdio.h>

// 有符号无符号。
int main(int argc, char *argv[]) {
  short a = 0100; // 八进制
  int b = -0x1; // 十六进制
  long c = 720; // 十进制

  unsigned short x = 0xffff; // 十六进制
  unsigned int y = 0x80000000; // 十六进制
  unsigned long z = 100; // 十进制

  puts("整数在写入内存之前可能会发生转换，在读取时也可能会发生转换，所以以下的输出会有些奇怪：");

  // 以无符号的形式输出有符号数
  printf("a = %#ho, b = %#x, c = %lu\n", a, b, c);

  // 以有符号数的形式输出无符号类型(只能以十进制形式输出)
  printf("x = %hd, b = %d, z = %ld\n", x, y, z);

  printf("\n");

  printf("不同进制的形式输出时对应的格式控制符(-表示没有对应的格式控制符):\n");
  printf("\t\tshort\tint\tlong\tunsigned short\tunsigned int\tunsigned long\n");
  printf("Otc\t\t-\t\t-\t-\t\t%%ho\t\t\t\t%%o\t\t\t\t%%lo\n");
  printf("Dec\t\t%%hd\t\t%%d\t%%ld\t\t%%hu\t\t\t\t%%u\t\t\t\t%%lu\n");
  printf("Hex\t\t-\t\t-\t-\t\t%%hx(X)\t\t\t%%x(X)\t\t\t%%lx(X)\n");
  return 0;
}
