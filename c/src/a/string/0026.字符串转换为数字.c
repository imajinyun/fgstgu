#include <stdio.h>
#include <stdlib.h>

#define LIMIT 30

char *str_gets(char *str, int n);

/**
 * 字符串转换为数字。
 *
 * 1. 当 base 分别为 10 和 16 时，字符串 "10" 分别被转换成数字 10 和 16；
 * 2. 如果 end 指向一个字符，*end 就是一个字符。因此，第 1 次转换在读到空字符时结束，
 *   此时 end 指向空字符。打印 end 会显示一个空字符串，以 %d 转换说明输出 *end 显示的是空字符的 ASCII 码；
 * 3. 许多实现使用 itoa() 和 ftoa() 函数分别把整数和浮点数转换成字符串。
 *   但是这两个函数并不是 C 标准库的成员，可以用 sprintf() 函数代替它们，因为 sprintf() 的兼容性更好；
 */
int main(int argc, char *argv[]) {
  char number[LIMIT];
  char *end;
  long value;
  puts("Enter a number(empty line to quit):");
  while (str_gets(number, LIMIT) && number[0] != '\0') {
    value = strtol(number, &end, 10);
    printf("base 10 input, base 10 output: %ld, stopped at %s (%d)\n", value, end, *end);
    value = strtol(number, &end, 16);
    printf("base 16 input, base 10 output: %ld, stopped at %s (%d)\n", value, end, *end);
    puts("Next number:");
  }
  puts("Done!\n");
  return 0;
}

char *str_gets(char *str, int n) {
  char *res;
  int i = 0;
  res = fgets(str, n, stdin);
  if (res) {
    while (str[i] != '\n' && str[i] != '\0') { i++; }
    if (str[i] == '\n') {
      str[i] = '\0';
    } else {
      while (getchar() != '\n') { continue; }
    }
  }
  return res;
}
