#include <stdio.h>

int main(int argc, char *argv[]) {
  int number;
  int sum = 0;
  int val;

  printf("请输入一个数字：");
  scanf("%d", &number);
  val = number;
  while (val) {
    sum = sum + val % 10;
    val /= 10;
    printf("sum = %d, val = %d\n", sum, val);
  }
  return 0;
}
