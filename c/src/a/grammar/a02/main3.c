#include <stdio.h>

/**
 * 英文字符。
 */
int main(int argc, char *argv[]) {
  puts("实例一：");

  char a = '1';
  char b = '@';
  char c = 'X';
  char d = ' ';

  // 使用 putchar() 输出
  putchar(a);
  putchar(' ');
  putchar(b);
  putchar(' ');
  putchar(c);
  putchar(' ');
  putchar(d);
  putchar(' ');
  putchar('\n');

  // 使用 printf() 输出
  printf("a = %c b = %c c = %c d = %d\n", a, b, c, d);

  puts("\n实例二：");

  int l = 'A';
  int m = 66;
  int n = 'C';

  printf("l = %c, %d\n", l, l);
  printf("m = %c, %d\n", m, m);
  printf("n = %c, %d\n", n, n);

  puts("\n实例三：");

  char str1[] = "I'm a first string.";
  char *str2 = "I'm a second string.";

  puts(str1);
  puts(str2);
  printf("%s\n%s\n", str1, str2);
  return 0;
}
