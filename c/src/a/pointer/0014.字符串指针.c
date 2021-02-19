#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
  char str[] = "https://www.apple.com/";
  char *ps = str;
  int len = strlen(str), i;

  printf("使用 *(ps + i)：");
  for (i = 0; i < len; i++) { printf("%c", *(ps + i)); }
  printf("\n");

  printf("使用 ps[i]：");
  for (i = 0; i < len; i++) { printf("%c", ps[i]); }
  printf("\n");

  printf("使用 *(str + i)：");
  for (i = 0; i < len; i++) { printf("%c", *(str + i)); }
  printf("\n");
  return 0;
}
