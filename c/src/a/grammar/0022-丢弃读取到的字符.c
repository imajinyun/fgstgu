#include <stdio.h>

int main(int argc, char *argv[]) {
  int a;
  char s[30];
  printf("请输入数字和字符串（包含大小写及数字）：\n");

  // %*d 表示读取一个整数并丢弃
  scanf("%*d %d", &a);

  // %*[a-z] 表示读取小写字母并丢弃
  scanf("%*[a-z]");

  // %*[^\n] 表示将换行符以外的字符全部丢弃
  scanf("%[^\n]", s);

  printf("a = %d, s = %s\n", a, s);
  return 0;
}
