#include <stdio.h>

// 进制表示。
int main(int argc, char *argv[]) {
  int a = 0b101, b = -0b11010, c = 0B100; // 二进制
  int d = 015, e = -0101, f = 017; // 八进制
  int g = 0x2A, h = -0xA0, i = 0XAA; // 十六进制

  printf("以十进制的形式输出：\n");
  printf("a = %d, b = %d, c = %d\n", a, b, c);
  printf("d = %d, e = %d, f = %d\n", d, e, f);
  printf("g = %d, h = %d, i = %d\n", g, h, i);

  printf("\n");

  short x = 0b101; // 二进制数字
  int y = 015; // 八进制数字
  long z = 0x2A; // 十六进制数字

  // 区分不同进制数字的一个简单办法：在输出时带上特定的前缀。即在格式控制符中加上 # 即可输出前缀。

  printf("以八进制的形式输出：\n");
  printf("x = %#ho, y = %#o, z = %#lo\n", x, y, z);

  printf("以十进制的形式输出：\n");
  printf("x = %hd, y = %d, z = %ld\n", x, y, z);

  printf("以十六进制小写的形式输出：\n");
  printf("x = %#hx, y = %#x, z = %#lx\n", x, y, z);

  printf("以十六进制大写的形式输出：\n");
  printf("x = %#hX, y = %#X, z = %#lX\n", x, y, z);
  return 0;
}
