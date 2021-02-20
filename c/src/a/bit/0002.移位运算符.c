#include <stdio.h>

int main(int argc, char *argv[]) {
  /**
   * 0b0010001010 <- 二进制数
   * 0b1000101000 <- 左移 2 位
   * 512 + 32 + 8 = 552
   */
  unsigned int v1 = 0b10001010;
  printf("左移：v1 = %d\n", v1 << 2);

  /**
   * 0b10001010 <- 二进制数
   * 0b00100010 <- 右移 2 位
   * 32 + 2 = 34
   */
  unsigned int v2 = 0b10001010;
  printf("右移：v2 = %d\n", v2 >> 2);
  return 0;
}
