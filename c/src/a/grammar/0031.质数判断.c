#include <stdio.h>

int main(int argc, char *argv[]) {
  int num;
  printf("输入一个数判断是否为质数：");
  scanf("%d", &num);

  int i;
  for (i = 2; i <= num; i++) {
    if (num % i == 0) { break; }
  }

  if (i == num) {
    printf("输入的 %d 是质数\n", num);
  } else {
    printf("输入的 %d 是合数\n", num);
  }
  return 0;
}
