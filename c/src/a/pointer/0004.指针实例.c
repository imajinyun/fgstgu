#include <stdio.h>

int main(int argc, char *argv[]) {
  char str[20] = "xyz.clang.com";

  char *s1 = str;
  char *s2 = str + 2;

  char c1 = str[4];
  char c2 = *str;
  char c3 = *(str + 4);
  char c4 = *str + 2;
  char c5 = (str + 1)[5];

  int num1 = *str + 2;
  long num2 = (long) str;     // 字符串 str 的首地址
  long num3 = (long) str + 2; // 字符串 str 的第 2 个元素的地址

  printf("s1 = %s\n", s1); // xyz.clang.com
  printf("s2 = %s\n", s2); // z.clang.com

  printf("\n");
  printf("c1 = %c\n", c1); // c
  printf("c2 = %c\n", c2); // x
  printf("c3 = %c\n", c3); // c
  printf("c4 = %c\n", c4); // z
  printf("c5 = %c\n", c5); // a

  printf("\n");
  printf("num1 = %d, num2 = %c\n", num1, num1); // num1 = 122, num2 = z
  printf("num2 = %ld\n", num2);                 // 140732824572848
  printf("num3 = %ld\n", num3);                 // 140732824572850
  return 0;
}
