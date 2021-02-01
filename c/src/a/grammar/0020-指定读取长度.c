#include <stdio.h>

void flush() {
  scanf("%*[^\n]");
  scanf("%*c");
}

int main(int argc, char *argv[]) {
  int a;
  float b;
  char c[30];

  printf("请输入变量 a b c 的值：\n");
  scanf("%2d", &a);
  flush();

  scanf("%5f", &b);
  flush();

  scanf("%22s", c);
  printf("a = %d b = %g c = %s\n", a, b, c);
  return 0;
}
