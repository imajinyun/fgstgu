#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
  char s1[] = "http://www.baidu.com/";
  char s2[] = "http://www.tmall.com/";

  int result = strcmp(s1, s2);
  if (result > 0) {
    printf("s1 > s2");
  } else if (result < 0) {
    printf("s1 < s2");
  } else {
    printf("s1 == s2");
  }
  printf("\n");
  return 0;
}
