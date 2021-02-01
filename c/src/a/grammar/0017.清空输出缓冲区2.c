#include <stdio.h>

int main(int argc, char *argv[]) {
  int a = 1, b = 2;
  char c;

  printf("请输入变量 a 和 b 的值：\n");
  scanf("a = %d", &a);
  while ((c = getchar()) != '\n' && c != EOF)
    ;

  scanf("b = %d", &b);
  printf("a = %d b = %d\n", a, b);
  return 0;
}
