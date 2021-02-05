#include <stdio.h>

int main(int argc, char *argv[]) {
  int i = 0;
  printf("请输入一个字符串：");
  while (getchar() != '\n') { i++; }
  printf("输入的字符串长度为：%d\n", i);
  return 0;
}
