#include <stdio.h>

int main(int argc, char *argv[]) {
  int i = 0, sum = 0;
  int number;
  printf("请输入一个数字：");
  scanf("%d", &number);
  while (i <= number) {
    sum += i;
    i++;
  }
  printf("从 1 + ... + %d 求和为：%d\n", number, sum);
  return 0;
}
