#include <stdio.h>

int main(int argc, char *argv[]) {
  int number;
  int sum = 0;
  int val;

  printf("请输入一个数字：");
  scanf("%d", &number);
  val = number;
  while (val) {
    sum = sum * 10 + val % 10;
    val /= 10;
    printf("sum = %d, val = %d\n", sum, val);
  }
  if (sum == number) {
    printf("输入的 %d 是回文数\n", number);
  } else {
    printf("输入的 %d 不是回文数\n", number);
  }
  return 0;
}
