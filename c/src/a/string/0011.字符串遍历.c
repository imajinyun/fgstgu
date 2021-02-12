#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
  char a[5];
  a[0] = 'a';
  a[1] = 'b';
  a[2] = 'c';
  a[3] = 'd';
  a[4] = 'e';

  char b[] = {"https://www.aliyun.com/"};
  char c[] = "https://cloud.tencent.com/";
  char d[] = {'hello world'};
  char e[5] = {0};

  puts("遍历字符串：");
  for (int i = 0; i < 5; i++) {
    printf("a[%d] = %3d", i, a[i]);
    printf("\tb[%d] = %c", i, b[i]);
    printf("\tc[%d] = %c", i, c[i]);
    printf("\td[%d] = %c", i, d[i]);
    printf("\te[%d] = %d", i, e[i]);
    printf("\n");
  }

  puts("\n字符串长度：");
  unsigned long arr[5] = {
    strlen(a), strlen(b), strlen(c), strlen(d), strlen(e),
  };
  for (int i = 0; i < 5; i++) {
    printf("%c = %lu ", a[i], arr[i]);
  }
  printf("\n");

  char str[30] = {0};
  char ch;
  int j;
  for (ch = 65, j = 0; ch <= 90; ch++, j++) { str[j] = ch; }
  printf("\n字符串：%s\n长度为：%lu\n", str, strlen(str));
  return 0;
}
