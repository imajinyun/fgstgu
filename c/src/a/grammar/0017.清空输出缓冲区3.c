#include <stdio.h>

int main(int argc, char *argv[]) {
  int a = 1, b = 2;

  printf("请输入变量 a 和 b：\n");
  scanf("a = %d", &a);
  scanf("%*[^\n]");
  scanf("%*c");
  scanf("b = %d", &b);
  printf("a = %d b = %d\n", a, b);
  return 0;
}
