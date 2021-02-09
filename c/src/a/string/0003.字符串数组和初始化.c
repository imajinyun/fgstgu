#include <stdio.h>

void test1() {
  int n = 8;
  char cookies[1];                        // 有效
  char cakes[2 + 5];                      // 有效，数组大小是整型常量表达式
  char pies[2 * sizeof(long double) + 1]; // 有效
  char crumbs[n];                         // 在 C99 标准之前无效，C99 标准之后这种数组是变长数组
}

void test2() {
  char car[10] = "Tesla";
  if (car == &car[0]) { printf("%s\n", car); }
  if (*car == 'T') { printf("%c\n", *car); };
  if (*(car + 1) == car[1] && car[1] == 'e') { printf("%c\n", car[1]); }
}

int main(int argc, char *argv[]) {
  const char a[40] = "Your time is limited, don't waste it";
  const char b[40] = {
    'Y', 'o', 'u', 'r', ' ', 't', 'i',  'm', 'e', ' ', 'i', 's', ' ', 'l', 'i', 'm', 'i', 't', 'e',
    'd', ',', ' ', 'd', 'o', 'n', '\'', 't', ' ', 'w', 'a', 's', 't', 'e', ' ', 'i', 't', '.', '\0',
  };
  const char c[] = "If you can't think of anything, fake it.";
  puts(a);
  puts(b);
  puts(c);

  test1();
  test2();
  return 0;
}
