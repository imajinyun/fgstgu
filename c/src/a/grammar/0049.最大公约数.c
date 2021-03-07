#include <stdio.h>

int commonDivisor(int a, int b);

int main(int argc, char *argv[]) {
  int a, b;
  printf("请数据两个数字：");
  scanf("%d %d", &a, &b);

  int result = commonDivisor(a, b);
  printf("最大公约数为：%d\n", result);
  return 0;
}

int commonDivisor(int a, int b) {
  if (a < b) {
    int t = a;
    a = b;
    b = t;
  }
  /**
   * 输入：2 9
   * 交换：9 2
   *
   * 第 1 次循环：t = 2; b = 9 % 2 = 1; a = 2;
   * 第 2 次循环：t = 1; b = 2 % 1 = 0; a = 1;
   */
  while (b != 0) {
    int t = b;
    b = a % b;
    a = t;
  }
  return a;
}
