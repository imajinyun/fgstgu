#include <ctype.h>
#include <stdio.h>
#include <string.h>

#define LIMIT 81

void toUpper(char *);
int punctCounter(const char *);

int main(int argc, char *argv[]) {
  // 该程序使用 fgets() 和 strchr() 组合，读取一行输入并把换行符替换成空字符。
  // 这种方法与使用 str_gets() 的区别是：str_gets() 会处理输入行剩余字符（如果有的话），为下一次输入做好准备。
  // 而本例只有一条输入语句，就没必要进行多余的步骤。
  char line[LIMIT];
  char *find;
  puts("Please enter a line:");
  fgets(line, LIMIT, stdin);
  find = strchr(line, '\n');
  if (find) { *find = '\0'; }
  toUpper(line);
  puts(line);
  printf("That line has %d punctuation characters\n", punctCounter(line));
  return 0;
}

void toUpper(char *str) {
  while (*str) {
    *str = toupper(*str);
    str++;
  }
}

int punctCounter(const char *str) {
  int count = 0;
  while (*str) {
    if (ispunct(*str)) { count++; }
    str++;
  }
  return count;
}
