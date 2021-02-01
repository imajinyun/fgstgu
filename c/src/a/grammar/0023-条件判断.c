#include <stdio.h>

int main(int argc, char *argv[]) {
  printf("请输入一个字符：");
  char ch;
  ch = getchar();

  if (ch < 32) {
    printf("这是一个控制字符\n");
  } else if (ch >= '0' && ch <= '9') {
    if (ch >= '5') {
      printf("这是一个大于等于 5 的数字\n");
    } else {
      printf("这是一个小于 5 的数字\n");
    }
  } else if (ch >= 'a' && ch <= 'z') {
    printf("这是小写字母\n");
  } else if (ch >= 'A' && ch <= 'Z') {
    printf("这是大写字母\n");
  } else {
    printf("其它字符\n");
  }
  return 0;
}
