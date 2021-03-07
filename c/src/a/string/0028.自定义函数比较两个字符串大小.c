#include <stdio.h>
#undef strcmp

int strcmp(char *s1, char *s2) {
  int i, result;

  for (i = 0; (result = s1[i] - s2[i]) == 0; i++) {
    if (s1[i] == '\0' || s2[i] == '\0') { break; }
  }
  return result;
}

int main(int argc, char *argv[]) {
  char s1[] = "http://www.baidu.com/";
  char s2[] = "http://www.sina.com.cn/";
  char s3[] = "http://www.qq.com/";

  int res3 = strcmp("Programmer", "Programming");
  printf("s1 - s2 = %d\n", strcmp(s1, s2));
  printf("s1 - s3 = %d\n", strcmp(s1, s3));
  printf("Programmer - Programming = %d\n", res3);
  return 0;
}
