#include <stdio.h>

int main(int argc, char *argv[]) {
  int number;

  printf("请输入一个数：");
  scanf("%d", &number);
  int f1, f2, f3;
  f1 = 1;
  f2 = 2;

  if (number == 1) {
    f3 = number;
  } else if (number == 2) {
    f3 = number;
  } else {
    for (int i = 3; i <= number; ++i) {
      f3 = f1 + f2;
      f1 = f2;
      f2 = f3;
    }
  }
  printf("索引为 %d 的斐波那契数列值为：%d\n", number, f3);
  return 0;
}
