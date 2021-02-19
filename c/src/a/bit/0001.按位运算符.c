#include <stdio.h>

int main(int argc, char *argv[]) {
  unsigned char v1 = 0b00000010;
  printf("按位取反（二进制反码）：\n");
  printf("~v1 = %d, v1 = %d\n\n", ~v1, v1); // ~v1 = 11111101 -> 取反加 1 -> 10000011 -> -3

  unsigned char v2 = (0b10010011) & (0b00111101); // 0b00010001 -> 17
  printf("按位与：\n");
  printf("v2 = %d\n\n", v2);

  unsigned char v3 = (0b10010011) | (0b00111101); // 0b10111111 -> 191
  printf("按位或：\n");
  printf("v3 = %d\n\n", v3);

  unsigned char v4 = (0b10010011) ^ (0b00111101); // 0b10101110 -> 174
  printf("按位异或：\n");
  printf("v4 = %d\n\n", v4);

  unsigned char opener = (0b00001111) | (0b10110110); // 10111111 -> 191
  printf("打开位（设置位）：");
  printf("setter = %d\n", opener);

  unsigned char closer = (0b00111101) & (0b10110110); // 00110100 -> 52
  printf("关闭位（清空位）：");
  printf("setter = %d\n", closer);
  return 0;
}
